package main

import (
	"goslug/configs/credentials"
	"goslug/internal/database"
	"goslug/internal/handlers"
	"goslug/internal/handlers/api"
	"goslug/internal/logger"
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

	logger.PrintMenu()

	// Initialize DB In Memory
	var memoryDB *database.MemoryStorage = &database.MemoryStorage{}
	err := memoryDB.Load()
	if err != nil {
		logger.Error("DATABASE Initialziation Dailed: " + err.Error())
		os.Exit(1)
	}

	logger.Log("DATABASE Initialziation Success")
	database.RedirectDB = memoryDB

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("website/public"))))
	logger.Log("SERVING_STATIC_FILES on /public")

	http.HandleFunc("/", handlers.Redirect)
	logger.Log("LISTENING on /", "METHODS: GET")

	http.HandleFunc("/login", handlers.Login)
	logger.Log("LISTENING on /login", " METHODS: GET, POST, DELETE")

	http.HandleFunc("/admin", middleware.RequireAuth(handlers.Admin))
	logger.Log("LISTENING on /admin", "METHODS: GET", "MIDDLEWARE: RequireAuth")

	http.HandleFunc("/api/redirects", middleware.RequireAuth(api.ApiRedirects))
	logger.Log("LISTENING on /api/redirects", "METHODS: POST, DELETE", "MIDDLEWARE: RequireAuth")

	logger.Log("SERVER_UP at http://" + HOST + PORT)
	http.ListenAndServe(HOST+PORT, nil)

}

func setEnv(variable string, def string) string {
	env, exist := os.LookupEnv(variable)
	if exist {
		logger.Log("ENV_FOUND " + variable + ":" + env)
		return env
	}

	logger.Log("ENV_NOT_FOUND::USING_DEFAULT " + variable + ":" + def)
	return def
}
