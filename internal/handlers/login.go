package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"goslug/configs/credentials"
	"goslug/configs/types"
	"html/template"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, req *http.Request) {

	authCookie, err := req.Cookie("goslug_session")

	switch req.Method {
	case "GET":
		// If im authenticated then just redirect to admin
		if err == nil && authCookie.Value == credentials.AUTH_KEY {
			http.Redirect(w, req, "/admin", http.StatusAccepted)
			return
		}

		//Otherwise return the login page
		tmpl, _ := template.ParseFiles("website/templates/login.html")
		tmpl.ExecuteTemplate(w, "login.html", nil)

	case "POST":
		var logIn types.LogInInfo
		err := json.NewDecoder(req.Body).Decode(&logIn)

		// Authenticated or not, i will give back a new auth cookie with 1 day of activity if credentials are rigth
		if err == nil {
			fmt.Print("username: " + logIn.Username + " password: " + logIn.Password)
			if logIn.Username == credentials.AUTH_USERNAME && logIn.Password == credentials.AUTH_PASSWORD {

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
				w.WriteHeader(http.StatusAccepted)

			} else {
				// IF we dont have valid credentials just repond with unautorized
				credentials.AUTH_KEY = "UNAUTORIZED"
				w.WriteHeader(http.StatusUnauthorized)
			}

		}
	case "DELETE":
		// If autenticated sent a empty cookie with no value to clear session (logout)
		if err == nil && authCookie.Value == credentials.AUTH_KEY {
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
			w.WriteHeader(http.StatusAccepted)

		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}
