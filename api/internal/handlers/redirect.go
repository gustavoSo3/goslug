package handlers

import (
	"fmt"
	"goslug/api/configs/types"
	"html/template"
	"net/http"
)

// /
func Redirect(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.NotFound(w, req)
		return
	}

	fmt.Print("Redirecting:\t\t", req.URL.Path)

	// If going to root just redirect to admin
	// TODO: Maybe later a welcome page?
	if req.URL.Path == "/" {
		tmpl, _ := template.ParseFiles("website/templates/home-page.html")
		const homePage string = "Home Page"
		tmpl.ExecuteTemplate(w, "home-page.html", homePage)
		return
	}

	// Search for rediredirect slug match
	for _, v := range types.SampleAdminData.Redirects {
		if v.Slug == req.URL.Path {
			fmt.Println("\t\t->\t\t", string(v.Destination))
			http.Redirect(w, req, string(v.Destination), http.StatusFound)
			return
		}
	}

	// If no matches at the moment return to admin or do a 404 bad request
	fmt.Println("\t\t->\t\t404 : Invalid Slug")
	http.Error(w, "Slug not in database", http.StatusNotFound)
}
