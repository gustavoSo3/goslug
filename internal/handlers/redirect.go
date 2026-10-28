package handlers

import (
	"goslug/internal/database"
	"goslug/internal/logger"
	"html/template"
	"net/http"
)

// Handler used to handle redirects, if not found redirects to welcome page
func Redirect(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {

		logger.Log(req.URL.Path + " " + req.Method + " not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Search for rediredirect slug match
	redirects, _ := database.RedirectDB.GetRedirects()
	for _, v := range redirects {
		if v.Slug == req.URL.Path {

			logger.Log(req.URL.Path + " -> " + v.Destination)
			http.Redirect(w, req, v.Destination, http.StatusSeeOther)
			return
		}
	}

	// If no matches at the moment return welcome page
	tmpl, _ := template.ParseFiles("website/templates/home-page.html")

	logger.Log(req.URL.Path + " Slug not database serving Welcome Page")
	tmpl.ExecuteTemplate(w, "home-page.html", nil)

}
