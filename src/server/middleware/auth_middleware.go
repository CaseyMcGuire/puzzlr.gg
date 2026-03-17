package middleware

import (
	"fmt"
	"net/http"

	"puzzlr.gg/src/server/reqctx"
	"puzzlr.gg/src/server/session"
)

// SessionMiddleware extracts the user ID from the session and adds it to the
// context if present, but does not require authentication.
func SessionMiddleware(sessionManager *session.SessionManager) (func(http.Handler) http.Handler, error) {
	if sessionManager == nil {
		return nil, fmt.Errorf("middleware.SessionMiddleware requires a non-nil sessionManager")
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, authenticated, err := sessionManager.AuthenticatedUserID(r)
			if err == nil && authenticated {
				ctx := reqctx.WithUserID(r.Context(), userID)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	}, nil
}

// AuthMiddleware requires authentication before advancing
func AuthMiddleware(sessionManager *session.SessionManager) (func(http.Handler) http.Handler, error) {
	if sessionManager == nil {
		return nil, fmt.Errorf("middleware.AuthMiddleware requires a non-nil sessionManager")
	}

	sessionMiddleware, err := SessionMiddleware(sessionManager)
	if err != nil {
		return nil, err
	}

	return func(next http.Handler) http.Handler {
		return sessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, err := reqctx.UserIDFromContext(r.Context()); err != nil {
				http.Error(w, "Unauthorized: please login", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}))
	}, nil
}
