package handlers

import (
	"goslug/configs/types"
	"goslug/internal/database"
	"html/template"
	"net/http"
)

// /admin
func Admin(w http.ResponseWriter, req *http.Request) {

	tmpl, _ := template.ParseFiles("website/templates/admin.html")

	redirects, _ := database.RedirectDB.GetRedirects()
	adminData := types.AdminData{
		Redirects: redirects,
	}

	tmpl.ExecuteTemplate(w, "admin.html", adminData)

}
