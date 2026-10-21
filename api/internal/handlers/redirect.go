package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"url-shortener/api/configs/types"
)

func RedirectHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.NotFound(w, req)
		return
	}

	fmt.Print("Redirecting:\t\t", req.URL.Path)

	if req.URL.Path == "/" {
		fmt.Println("\t\t->\t\t/admin : root")
		http.Redirect(w, req, "/admin", http.StatusFound)
		return
	}

	for _, v := range types.SampleAdminData.Redirects {
		if v.Slug == template.URL(req.URL.Path) {
			fmt.Println("\t\t->\t\t", string(v.Destination))
			http.Redirect(w, req, string(v.Destination), http.StatusFound)
			return
		}
	}

	fmt.Println("\t\t->\t\t/admin : Invalid Slug")
	http.Redirect(w, req, "/admin", http.StatusFound)
}
