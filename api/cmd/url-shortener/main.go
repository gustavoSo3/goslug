package main

import (
	"fmt"
	"net/http"
	"url-shortener/api/internal/handlers"
)

func main() {
	printMenu()

	http.HandleFunc("/", handlers.RedirectHandler)
	fmt.Println("listening\t/* and redirecting")

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("website/public"))))
	fmt.Println("setting\t\t/public")

	http.HandleFunc("/admin", handlers.AdminHandler)
	fmt.Println("listening\t/admin")

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
