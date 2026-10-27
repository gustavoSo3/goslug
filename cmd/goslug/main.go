package main

import (
	"fmt"
	"goslug/configs/credentials"
	"goslug/internal/database"
	"goslug/internal/handlers"
	"goslug/internal/handlers/api"
	"goslug/internal/middleware"
	"net/http"
	"os"
)

func main() {

	// GET ENVs
	var HOST string
	var PORT string

	PORT = setEnv("PORT", ":8080")
	HOST = setEnv("HOST", "0.0.0.0")

	credentials.AUTH_USERNAME = setEnv("USERNAME", "admin")
	credentials.AUTH_PASSWORD = setEnv("PASSWORD", "password")

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

	http.HandleFunc("/login", handlers.Login)

	http.HandleFunc("/admin", middleware.RequireAuth(handlers.Admin))
	fmt.Println("listening\t/admin", "GET, POST, DELETE")

	http.HandleFunc("/api/redirects", middleware.RequireAuth(api.ApiRedirects))
	fmt.Println("listening\t/api/redirects ", "POST, DELETE")

	fmt.Println("Server up at: http://" + HOST + PORT + "\n\n")
	http.ListenAndServe(HOST+PORT, nil)

}

func setEnv(variable string, def string) string {
	env, exist := os.LookupEnv(variable)
	if exist {
		return env
	}
	return def
}

func printMenu() {
	fmt.Print("===========================================\n")
	fmt.Print("GoSlug\n")
	fmt.Print("Wellcome\n")
	fmt.Print("https://github.com/gustavoSo3/goslug\n")
	fmt.Print("v0.1\n")
	fmt.Print("===========================================\n")
}
