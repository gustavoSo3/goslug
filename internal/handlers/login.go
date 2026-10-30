package handlers

import (
	"encoding/json"
	"goslug/configs/credentials"
	"goslug/configs/types"
	"goslug/internal/logger"
	"html/template"
	"net/http"
	"time"
)

// Handler used to manage Login different paths if user is authenticated or not
func Login(w http.ResponseWriter, req *http.Request) {

	authCookie, err := req.Cookie("goslug_session")

	var authenticated bool = err == nil && credentials.ValidateAuthentication(authCookie.Value)

	switch req.Method {
	case "GET":

		// If im authenticated then just redirect to admin
		if authenticated {
			logger.Log("GET Already authenticated redirecting to /admin")
			http.Redirect(w, req, "/admin", http.StatusSeeOther)
			return
		}

		//Otherwise return the login page
		tmpl, _ := template.ParseFiles("website/templates/login.html")

		logger.Log("Not authenticated serving Log In Page")
		tmpl.ExecuteTemplate(w, "login.html", nil)

	case "POST":

		// If im authenticated then just redirect to admin
		if authenticated {
			logger.Log("POST Already authenticated redirecting to /admin")
			http.Redirect(w, req, "/admin", http.StatusSeeOther)
			return
		}

		// Decode request body
		var login types.LoginInfo
		err := json.NewDecoder(req.Body).Decode(&login)
		if err != nil {
			return
		}

		logger.Log("AUTH_ATTEMPT", "username:", login.Username)
		authKey, successfulAuthentication := credentials.Authenticate(login.Username, login.Password)

		if successfulAuthentication {

			authCookie = &http.Cookie{
				Name:     "goslug_session",
				Value:    authKey,
				Path:     "/",
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}

			http.SetCookie(w, authCookie)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}

	case "DELETE":
		// If not autenticated sent a empty cookie with no value to clear session (logout)
		if !authenticated {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		logger.Log("AUTH_LOGGED_OUT")

		authCookie = &http.Cookie{
			Name:     "goslug_session",
			Value:    "",
			Path:     "/",
			Expires:  time.Now(),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, authCookie)
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
