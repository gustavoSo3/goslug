package main

import (
	"goslug/configs/credentials"
	"goslug/internal/database"
	"goslug/internal/handlers"
	"goslug/internal/handlers/api"
	"goslug/internal/logger"
	"goslug/internal/middleware"
	"goslug/internal/utils"
	"net/http"
	"os"
)

func main() {

	// GET ENVs
	var HOST string = utils.SetEnv("HOST", "localhost")
	var PORT string = utils.SetEnv("PORT", "8080")
	var STORAGE_DRIVER string = utils.SetEnv("STORAGE_DRIVER", "json")

	credentials.AUTH_USERNAME = utils.SetEnv("USERNAME", "admin")
	credentials.AUTH_PASSWORD = utils.SetEnv("PASSWORD", "password")

	logger.PrintMenu()

	var err error
	switch STORAGE_DRIVER {
	case "memory":
		// Initialize DB In Memory
		var memoryDB *database.MemoryDB = &database.MemoryDB{}
		err = memoryDB.Load()
		if err != nil {
			// If error exit since there is no alternative for memory
			logger.Error("MEMORY_DB Initialziation failed: " + err.Error())
			os.Exit(1)
		}
		logger.Log("DATABASE MEMORY_DB Initialized")
		database.RedirectDB = memoryDB
	case "json":
		var jsonDB *database.JSONDB = &database.JSONDB{}
		err = jsonDB.Load()
		if err != nil {
			logger.Error("JSON_DB Initialization failed: " + err.Error())
			os.Exit(1)
		}
		logger.Log("DATABASE JSON_DB Initialized")
		database.RedirectDB = jsonDB
	default:
		logger.Error("INVALID_ENV STORAGE_DRIVER ", STORAGE_DRIVER)
		os.Exit(1)
	}

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

	logger.Log("SERVER_UP at http://" + HOST + ":" + PORT)
	http.ListenAndServe(HOST+":"+PORT, nil)

}
