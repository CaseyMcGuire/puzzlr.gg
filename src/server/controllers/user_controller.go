package controllers

import (
	"fmt"
	"net/http"

	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/services"
	"puzzlr.gg/src/server/session"
	"puzzlr.gg/src/server/views"
)

type UserController struct {
	userService    *services.UserService
	sessionManager *session.SessionManager
}

func (u *UserController) HandleRegisterGet(w http.ResponseWriter, r *http.Request) {
	if u.sessionManager.RedirectIfAuthenticated(w, r, "/") {
		return
	}
	err := views.ReactPage("Register", "index").Render(w)
	if err != nil {
		http.Error(w, "failed to render page", http.StatusInternalServerError)
	}
}

func (u *UserController) HandleRegisterPost(w http.ResponseWriter, r *http.Request) {
	if u.sessionManager.RedirectIfAuthenticated(w, r, "/") {
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	exists, err := u.userService.UserExists(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
}

func NewUserController(
	dbClient *ent.Client,
	sessionManager *session.SessionManager,
) (*UserController, error) {
	if sessionManager == nil {
		return nil, fmt.Errorf("controllers.NewUserController requires a non-nil sessionManager")
	}

	userService, err := services.NewUserService(dbClient)
	if err != nil {
		return nil, err
	}

	return &UserController{
		userService:    userService,
		sessionManager: sessionManager,
	}, nil
}
