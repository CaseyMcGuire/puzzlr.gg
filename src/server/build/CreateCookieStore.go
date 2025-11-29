package build

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
)

func CreateCookieStore() *sessions.CookieStore {
	authKey := []byte(os.Getenv("SESSION_AUTH_KEY"))             // For authentication (HMAC)
	encryptionKey := []byte(os.Getenv("SESSION_ENCRYPTION_KEY")) // For encryption (AES)

	if len(authKey) == 0 || len(encryptionKey) == 0 {
		if os.Getenv("APP_ENV") == "production" {
			log.Fatalf("Authentication and/or encryption key(s) are not set")
		}
		// Fallback for local development if env vars are not set
		authKey = securecookie.GenerateRandomKey(64)
		encryptionKey = securecookie.GenerateRandomKey(32)
	}

	store := sessions.NewCookieStore(
		authKey,
		encryptionKey,
	)

	store.Options = &sessions.Options{
		Path:     "/",       // Apply to whole site
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,      // Prevent JavaScript access to the cookie
		Secure:   true,      // Only send cookie over HTTPS
		SameSite: http.SameSiteLaxMode,
	}
	// For local development without HTTPS
	if os.Getenv("APP_ENV") != "production" {
		store.Options.Secure = false
	}
	return store
}
