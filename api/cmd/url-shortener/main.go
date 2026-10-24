package main

import (
	"fmt"
	"net/http"
	"url-shortener/api/internal/handlers"
	"url-shortener/api/internal/handlers/api"
)

func main() {
	printMenu()

	http.HandleFunc("/", handlers.Redirect)
	fmt.Println("listening\t/* and redirecting")

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("website/public"))))
	fmt.Println("setting\t\t/public")

	http.HandleFunc("/admin", handlers.Admin)
	fmt.Println("listening\t/admin")

	http.HandleFunc("/api/redirects", api.ApiRedirects)
	fmt.Println("listening\t/api/create-redirect")

	fmt.Println("Server up at: http://localhost:8080\n\n")
	http.ListenAndServe(":8080", nil)

}

func printMenu() {
	fmt.Print("===========================================\n")
	fmt.Print("URL-SHORTENER\n")
	fmt.Print("Wellcome\n")
	fmt.Print("https://github.com/gustavoSo3/url-shortener\n")
	fmt.Print("v0.1\n")
	fmt.Print("===========================================\n")
}
