package handlers

import (
	"crypto/sha256"
	"encoding/hex"
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

	var authenticated bool = false
	authCookie, err := req.Cookie("goslug_session")
	if err == nil && credentials.ValidAuthKey(authCookie.Value) {
		authenticated = true
	}

	switch req.Method {
	case "GET":

		// If im authenticated then just redirect to admin
		if authenticated {
			logger.Log("Already authenticated redirecting to /admin")
			http.Redirect(w, req, "/admin", http.StatusSeeOther)
			return
		}

		//Otherwise return the login page
		tmpl, _ := template.ParseFiles("website/templates/login.html")

		logger.Log("Not authenticated serving Log In Page")
		tmpl.ExecuteTemplate(w, "login.html", nil)

	case "POST":

		// Decode request body
		var logIn types.LogInInfo
		err := json.NewDecoder(req.Body).Decode(&logIn)
		if err != nil {
			return
		}

		logger.Log("AUTH_ATTEMPT", "username:", logIn.Username, "password:", logIn.Password)

		if !credentials.ValidLogInCredentials(logIn.Username, logIn.Password) {
			// IF we dont have valid credentials just repond with unautorized
			logger.Log("AUTH_ATTEMPT", "FAILED Invalid credentials")
			credentials.AUTH_KEY = "UNAUTORIZED"
			w.WriteHeader(http.StatusUnauthorized)
			return

		}

		logger.Log("AUTH_ATTEMPT", "Valid credentials")
		hash := sha256.New()
		hash.Write([]byte(string(credentials.AUTH_USERNAME + credentials.AUTH_PASSWORD + time.Now().String())))
		credentials.AUTH_KEY = hex.EncodeToString(hash.Sum(nil))

		authCookie = &http.Cookie{
			Name:     "goslug_session",
			Value:    credentials.AUTH_KEY,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, authCookie)
		w.WriteHeader(http.StatusCreated)

	case "DELETE":
		// If autenticated sent a empty cookie with no value to clear session (logout)
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
