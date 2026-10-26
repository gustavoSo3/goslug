package main

import (
	"fmt"
	"goslug/api/internal/database"
	"goslug/api/internal/handlers"
	"goslug/api/internal/handlers/api"
	"net/http"
	"os"
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

	// Initialize DB In Memory
	var memoryDB *database.MemoryStorage = &database.MemoryStorage{}
	memoryDB.Load()
	database.RedirectDB = memoryDB

	// Set endpoints
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
