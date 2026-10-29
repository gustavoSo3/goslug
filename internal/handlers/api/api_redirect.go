package api

import (
	"encoding/json"
	"goslug/configs/enums"
	"goslug/configs/types"
	"goslug/internal/database"
	"goslug/internal/logger"
	"net/http"
)

// Handler used for calls to /api/redirects
func ApiRedirects(w http.ResponseWriter, req *http.Request) {

	// If its POST
	switch req.Method {
	case "POST":

		// Decoding body
		var newRedirect types.Redirect
		json.NewDecoder(req.Body).Decode(&newRedirect)
		logger.Log("NEW_REDIRECT ", newRedirect.Slug, " ", newRedirect.Destination, " ", newRedirect.Description)

		//Add Redirect
		err := database.RedirectDB.AddRedirect(newRedirect)
		if err != nil {
			switch err.Error() {
			case enums.ERRORExistingSlug:
				// Bad Redirect Data
				logger.Log("NEW_REDIRECT Not created", enums.ERRORExistingSlug)
				http.Error(w, enums.ERRORExistingSlug, http.StatusBadRequest)

			case enums.ERRORBadRedirectData:
				// Bad Redirect Data
				logger.Log("NEW_REDIRECT Not created", enums.ERRORBadRedirectData)
				http.Error(w, enums.ERRORBadRedirectData, http.StatusBadRequest)
			default:
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}
		database.RedirectDB.Save()
		logger.Log("NEW_REDIRECT Created")
		w.WriteHeader(http.StatusCreated)

	case "DELETE":
		// Decoding body
		var slugToDelete types.Slug
		json.NewDecoder(req.Body).Decode(&slugToDelete)
		logger.Log("DELETING_SLUG ", slugToDelete)

		err := database.RedirectDB.RemoveRedirect(slugToDelete)
		if err != nil {
			// If Slug not found
			logger.Log("DELETING_SLUG ", enums.ERRORNoMatchingSlug)
			http.Error(w, enums.ERRORNoMatchingSlug, http.StatusBadRequest)
			return
		}

		logger.Log("DELETING_SLUG Deleted")
		database.RedirectDB.Save()
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
