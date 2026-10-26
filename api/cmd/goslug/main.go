package main

import (
	"fmt"
	"net/http"
	"os"
	"url-shortener/api/internal/handlers"
	"url-shortener/api/internal/handlers/api"
)

func main() {

	// GET ENVs
	var HOST string
	var PORT string

	envPort, exist := os.LookupEnv("PORT")
	if exist {
		PORT = envPort
	} else {
		PORT = ":8080"
	}

	envAddress, exist := os.LookupEnv("HOST")
	if exist {
		HOST = envAddress
	} else {
		HOST = "0.0.0.0"
	}

	printMenu()

	http.HandleFunc("/", handlers.Redirect)
	fmt.Println("listening\t/ and redirecting")

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("website/public"))))
	fmt.Println("setting\t\t/public")

	http.HandleFunc("/admin", handlers.Admin)
	fmt.Println("listening\t/admin")

	http.HandleFunc("/api/redirects", api.ApiRedirects)
	fmt.Println("listening\t/api/redirects ", "POST, DELETE")

	fmt.Println("Server up at: http://" + HOST + PORT + "\n\n")
	http.ListenAndServe(HOST+PORT, nil)

}

func printMenu() {
	fmt.Print("===========================================\n")
	fmt.Print("GoSlug\n")
	fmt.Print("Wellcome\n")
	fmt.Print("https://github.com/gustavoSo3/goslug\n")
	fmt.Print("v0.1\n")
	fmt.Print("===========================================\n")
}
