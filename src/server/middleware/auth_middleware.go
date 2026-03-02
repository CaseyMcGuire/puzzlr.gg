package middleware

import (
	"net/http"

	"puzzlr.gg/src/server/build"
	"puzzlr.gg/src/server/controllers"
	"puzzlr.gg/src/server/reqctx"
)

// SessionMiddleware extracts the user ID from the session and adds it to the
// context if present, but does not require authentication.
func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := build.CreateCookieStore().Get(r, controllers.SessionName)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		authenticated, okAuth := session.Values[controllers.Authenticated].(bool)
		userID, okUserID := session.Values[controllers.UserID].(int)

		if okAuth && authenticated && okUserID && userID != -1 {
			ctx := reqctx.WithUserID(r.Context(), userID)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware requires authentication before advancing
func AuthMiddleware(next http.Handler) http.Handler {
	return SessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := reqctx.UserIDFromContext(r.Context()); err != nil {
			http.Error(w, "Unauthorized: please login", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}))
}
