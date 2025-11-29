package middleware

import (
	"context"
	"net/http"

	"puzzlr.gg/src/server/build"
	"puzzlr.gg/src/server/controllers"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store := build.CreateCookieStore()
		session, err := store.Get(r, controllers.Authenticated)
		if err != nil {
			// If the cookie is invalid or not present, store.Get might still return
			// a new empty session, so checking err is important for store issues.
			// However, the primary check is for the authenticated flag or userID.
			http.Error(w, "Unauthorized: session error", http.StatusUnauthorized)
			return
		}

		// Check if the user is authenticated
		auth, okAuth := session.Values[controllers.Authenticated].(bool)
		userID, okUserID := session.Values[controllers.UserID].(int)

		if !okAuth || !auth || !okUserID || userID == -1 {
			http.Error(w, "Unauthorized: please login", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))

		// For simplicity here, we just proceed.
		// In a real app, you'd likely want to pass the userID along.
		// log.Printf("User %s authenticated for %s", userID, r.URL.Path)
	})
}
