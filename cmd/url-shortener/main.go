package main

import (
	"fmt"
	"net/http"
)

func main() {
	printMenu()

	http.HandleFunc("/", rootHandler)
	fmt.Println("listening\t/admin")

	http.HandleFunc("/admin", adminHandler)
	fmt.Println("listening\t/admin")

	fmt.Println("Server up at: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/admin", http.StatusFound)
}

func adminHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ADMIN")
}
func printMenu() {
	fmt.Print("===========================================\n")
	fmt.Print("URL-SHORTENER\n")
	fmt.Print("Wellcome\n")
	fmt.Print("https://github.com/gustavoSo3/url-shortener\n")
	fmt.Print("v0.1\n")
	fmt.Print("===========================================\n")
}
