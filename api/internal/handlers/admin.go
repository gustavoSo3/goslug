package handlers

import (
	"goslug/api/configs/types"
	"html/template"
	"net/http"
)

// /admin
func Admin(w http.ResponseWriter, req *http.Request) {

	tmpl, _ := template.ParseFiles("website/templates/admin.html")

	adminData := types.SampleAdminData

	tmpl.ExecuteTemplate(w, "admin.html", adminData)

}
