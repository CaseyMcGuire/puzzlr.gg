package controllers

import (
	ent "gameboard/src/server/db/ent/codegen"
	"gameboard/src/server/services"
	"gameboard/src/server/views"
	"github.com/gorilla/sessions"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

var (
	store         = sessions.NewCookieStore([]byte("secret-key"))
	authenticated = "authenticated"
)

func (u *UserController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user_session")
	if auth, ok := session.Values[authenticated].(bool); ok && auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		views.ReactPage("Login", "index").Render(w)
		return
	} else if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		exists, err := u.userService.UserWithPasswordExists(r.Context(), email, password)
		if err != nil {
			http.Redirect(w, r, "/500", http.StatusInternalServerError)
			return
		}

		if exists {
			session.Values[authenticated] = true
			sessionErr := session.Save(r, w)
			if sessionErr == nil {
				http.Redirect(w, r, "/", http.StatusOK)
			} else {
				http.Error(w, sessionErr.Error(), http.StatusInternalServerError)
			}
		} else {
			http.Redirect(w, r, "/login?no_such_user=true", http.StatusFound)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (u *UserController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user_session")
	if auth, ok := session.Values[authenticated].(bool); ok && auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		views.ReactPage("Register", "index").Render(w)
		return
	} else if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		exists, err := u.userService.UserExists(r.Context(), email)
		if err != nil {
			http.Redirect(w, r, "/500", http.StatusInternalServerError)
		} else if exists {
			http.Redirect(w, r, "/register?error=username_already_taken", http.StatusSeeOther)
		} else {
			_, err := u.userService.CreateUser(r.Context(), email, password)
			if err != nil {
				http.Error(w, "Something went wrong. Please try again", http.StatusInternalServerError)
			} else {
				http.Redirect(w, r, "/login", http.StatusFound)
			}
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (u *UserController) HandleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user_session")

	// Set MaxAge to -1 to delete the cookie
	session.Options.MaxAge = -1
	_ = session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func NewUserController(
	dbClient ent.Client,
) *UserController {
	return &UserController{
		userService: *services.NewUserService(
			dbClient,
		),
	}
}
