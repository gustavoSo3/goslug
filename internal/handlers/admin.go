package handlers

import (
	"goslug/internal/database"
	"goslug/internal/logger"
	"html/template"
	"net/http"
)

// Handler for admin page
func Admin(w http.ResponseWriter, req *http.Request) {

	tmpl, _ := template.ParseFiles("website/templates/admin.html")

	redirects, _ := database.RedirectDB.GetRedirects()

	logger.Log(req.URL.Path + " Serving Admin Console")
	tmpl.ExecuteTemplate(w, "admin.html", redirects)

}
