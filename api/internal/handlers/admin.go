package handlers

import (
	"html/template"
	"net/http"
	"url-shortener/api/configs/types"
)

// /admin
func Admin(w http.ResponseWriter, req *http.Request) {

	tmpl, _ := template.ParseFiles("website/admin.html")

	adminData := types.SampleAdminData

	if types.SampleAdminData.Host == "" {
		types.SampleAdminData.Host = "https://short.sotres.dev"
	}

	tmpl.ExecuteTemplate(w, "admin.html", adminData)

}
