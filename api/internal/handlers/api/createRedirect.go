package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"url-shortener/api/configs/types"
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

		if newRedirect.Description != "" && newRedirect.Slug != "" && newRedirect.Destination != "" {
			fmt.Print("\tValid Redirect\t")
			if isSlugUnique(newRedirect.Slug) {
				fmt.Print("\tUnique Slug\t")

				// add it to the storage. ATP using just local mememory
				// TODO: This needs to be done in the DB
				types.SampleAdminData.Redirects = append(types.SampleAdminData.Redirects, newRedirect)

				fmt.Println("Redirect Saved")

				// Return confirmation of object reated if success
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(newRedirect)
				return

			} else {
				// Slug already exist
				fmt.Println("\tREQ:ERR SLUG ALREADY EXIST\t")
				http.Error(w, "Slug already exist", http.StatusBadRequest)
			}

		} else {
			// Bad Redirect Data
			fmt.Println("\tREQ:ERR BAD REDIRECT\t")
			http.Error(w, "Bad Redirect Data", http.StatusBadRequest)
		}

	//IF its DELETE
	case "DELETE":
		// Decoding body
		var slugToDelete string
		json.NewDecoder(req.Body).Decode(&slugToDelete)
		fmt.Print("Dele Slug: ", slugToDelete)

		// Looking for matches
		for i, v := range types.SampleAdminData.Redirects {
			if slugToDelete == v.Slug {
				// If found delete and return the slug deleted
				// TODO: This needs to be done in the DB
				types.SampleAdminData.Redirects = slices.Delete(types.SampleAdminData.Redirects, i, i+1)

				fmt.Println("\tDELETED")

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, slugToDelete)
				return
			}
		}

		// If Lug not found
		fmt.Println("\tREQ:ERR SLUG not found\t")
		http.Error(w, "Slug not found", http.StatusBadRequest)
		http.NotFound(w, req)
	}

}

func isSlugUnique(slug string) bool {
	for _, v := range types.SampleAdminData.Redirects {
		if slug == v.Slug {
			return false
		}
	}
	return true
}
