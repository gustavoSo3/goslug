package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/api/configs/types"
)

// /api/redirects
func Redirect(w http.ResponseWriter, req *http.Request) {

	// If its POST
	if req.Method == "POST" {

		// Decoding body
		var newRedirect types.Redirect
		json.NewDecoder(req.Body).Decode(&newRedirect)
		fmt.Print("New Redirect: ", newRedirect)

		if newRedirect.Description != "" && newRedirect.Slug != "" && newRedirect.Destination != "" {
			fmt.Print("\tValid Redirect\t")
			if isSlugUnique(newRedirect.Slug) {
				fmt.Print("\tUnique Slug\t")

				// add it to the storage. ATP using just local mememory, this will be a write to the db later
				types.SampleAdminData.Redirects = append(types.SampleAdminData.Redirects, newRedirect)

				fmt.Println("Redirect Saved")

				// Return confirmation of object reated if success
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)

				json.NewEncoder(w).Encode(newRedirect)
			} else {
				// Slug already exist
				fmt.Println("\tREQ:ERR SLUG ALREADY EXIST\t")

				http.NotFound(w, req)
			}

		} else {
			// Incomplete Redirect
			fmt.Println("\tREQ:ERR BAD REDIRECT\t")
			http.NotFound(w, req)
		}

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
