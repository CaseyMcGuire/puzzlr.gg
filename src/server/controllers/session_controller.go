package controllers

import (
	"fmt"
	"net/http"

	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/services"
	"puzzlr.gg/src/server/session"
	"puzzlr.gg/src/server/views"
)

type SessionController struct {
	userService    *services.UserService
	sessionManager *session.SessionManager
}

func (sc *SessionController) HandleLoginGet(w http.ResponseWriter, r *http.Request) {
	if sc.sessionManager.RedirectIfAuthenticated(w, r, "/") {
		return
	}
	err := views.ReactPage("Login", "index").Render(w)
	if err != nil {
		http.Error(w, "failed to render page", http.StatusInternalServerError)
	}
}

func (sc *SessionController) HandleLoginPost(w http.ResponseWriter, r *http.Request) {
	if sc.sessionManager.RedirectIfAuthenticated(w, r, "/") {
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := sc.userService.GetUserWithPassword(r.Context(), email, password)
	if err != nil {
		http.Error(w, "failed to authenticate user", http.StatusInternalServerError)
		return
	}

	if user != nil {
		if err := sc.sessionManager.SignIn(w, r, user.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login?no_such_user=true", http.StatusFound)
	}
}

func (sc *SessionController) HandleLogout(w http.ResponseWriter, r *http.Request) {
	if err := sc.sessionManager.SignOut(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func NewSessionController(
	dbClient *ent.Client,
	sessionManager *session.SessionManager,
) (*SessionController, error) {
	if sessionManager == nil {
		return nil, fmt.Errorf("controllers.NewSessionController requires a non-nil sessionManager")
	}

	userService, err := services.NewUserService(dbClient)
	if err != nil {
		return nil, err
	}

	return &SessionController{
		userService:    userService,
		sessionManager: sessionManager,
	}, nil
}
