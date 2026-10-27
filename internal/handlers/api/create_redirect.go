package api

import (
	"encoding/json"
	"fmt"
	"goslug/configs/enums"
	"goslug/configs/types"
	"goslug/internal/database"
	"net/http"
)

// /api/redirects
func ApiRedirects(w http.ResponseWriter, req *http.Request) {

	// If its POST
	switch req.Method {
	case "POST":
		// Decoding body
		var newRedirect types.Redirect
		json.NewDecoder(req.Body).Decode(&newRedirect)
		fmt.Print("New Redirect: ", newRedirect)

		//Add Redirect
		err := database.RedirectDB.AddRedirect(newRedirect)
		if err != nil {
			switch err.Error() {

			case enums.ERRORExistingSlug:
				// Bad Redirect Data
				fmt.Println("\tREQ:ERR Slug already exist\t")
				http.Error(w, "Slug in use", http.StatusBadRequest)

			case enums.ERRORBadRedirectData:
				// Bad Redirect Data
				fmt.Println("\tREQ:ERR BAD REDIRECT\t")
				http.Error(w, "Bad Redirect Data", http.StatusBadRequest)

			}
			return
		}

		// Return confirmation of object reated if success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newRedirect)

	//IF its DELETE
	case "DELETE":
		// Decoding body
		var slugToDelete string
		json.NewDecoder(req.Body).Decode(&slugToDelete)
		fmt.Print("Dele Slug: ", slugToDelete)

		err := database.RedirectDB.RemoveRedirect(slugToDelete)
		if err != nil {
			// If Slug not found
			fmt.Println("\tREQ:ERR SLUG not found\t")
			http.Error(w, "Slug not found", http.StatusBadRequest)
			http.NotFound(w, req)

			return
		}

		fmt.Println("\tDELETED")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, slugToDelete)

	}

}
