package session

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const (
	sessionName      = "user_session"
	authenticatedKey = "authenticated"
	userIDKey        = "UserID"
)

type SessionManager struct {
	store *sessions.CookieStore
}

func NewSessionManager(store *sessions.CookieStore) (*SessionManager, error) {
	if store == nil {
		return nil, fmt.Errorf("session manager requires a non-nil cookie store")
	}

	return &SessionManager{
		store: store,
	}, nil
}

func (m *SessionManager) RedirectIfAuthenticated(w http.ResponseWriter, r *http.Request, redirectTo string) bool {
	_, authenticated, err := m.AuthenticatedUserID(r)
	if err != nil || !authenticated {
		return false
	}

	http.Redirect(w, r, redirectTo, http.StatusSeeOther)
	return true
}

func (m *SessionManager) AuthenticatedUserID(r *http.Request) (userID int, authenticated bool, err error) {
	session, err := m.store.Get(r, sessionName)
	if err != nil {
		if isSessionDecodeError(err) {
			return 0, false, nil
		}
		return 0, false, err
	}

	authenticated, okAuth := session.Values[authenticatedKey].(bool)
	userID, okUserID := session.Values[userIDKey].(int)

	if !okAuth || !authenticated || !okUserID || userID <= 0 {
		return 0, false, nil
	}

	return userID, true, nil
}

func (m *SessionManager) SignIn(w http.ResponseWriter, r *http.Request, userID int) error {
	session, err := m.store.Get(r, sessionName)
	if err != nil && !isSessionDecodeError(err) {
		return err
	}

	session.Values[authenticatedKey] = true
	session.Values[userIDKey] = userID

	return session.Save(r, w)
}

func (m *SessionManager) SignOut(w http.ResponseWriter, r *http.Request) error {
	session, err := m.store.Get(r, sessionName)
	if err != nil && !isSessionDecodeError(err) {
		return err
	}

	session.Values[authenticatedKey] = false
	session.Values[userIDKey] = -1
	session.Options.MaxAge = -1

	return session.Save(r, w)
}

func isSessionDecodeError(err error) bool {
	var secureCookieErr securecookie.Error
	return errors.As(err, &secureCookieErr) && secureCookieErr.IsDecode()
}
