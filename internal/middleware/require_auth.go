package middleware

import (
	"goslug/configs/credentials"
	"net/http"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		authCookie, err := req.Cookie("goslug_session")

		if err != nil || authCookie.Value != credentials.AUTH_KEY {
			// Not authenticated redirect to /login
			http.Redirect(w, req, "/login", http.StatusSeeOther)
			return
		}

		// Uthenticated
		next(w, req)
	}
}
