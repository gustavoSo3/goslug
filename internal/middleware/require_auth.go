package middleware

import (
	"goslug/configs/credentials"
	"goslug/internal/logger"
	"net/http"
)

// Middleware to handle endpoints that need authentication
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		authCookie, err := req.Cookie("goslug_session")

		if err != nil || !credentials.ValidateAuthentication(authCookie.Value) {
			// Not authenticated redirect to /login
			logger.Log("=====Unauthenticated -> /login =====")
			http.Redirect(w, req, "/login", http.StatusSeeOther)
			return
		}

		// Uthenticated
		logger.Log("=====Authenticated=====")
		next(w, req)
	}
}
