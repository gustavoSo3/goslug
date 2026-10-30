package database

import (
	"encoding/json"
	"errors"
	"goslug/configs/enums"
	"goslug/configs/types"
	"goslug/internal/logger"
	"goslug/internal/utils"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

type JSONDB struct {
	Redirects    map[types.Slug]types.Redirect `json:"redirects"`
	jsonFilepath string
	mu           sync.Mutex
}

func (db *JSONDB) Load() error {
	db.mu.Lock()
	defer db.mu.Unlock()
	// Set default path for storage file
	if db.jsonFilepath == "" {
		db.jsonFilepath = "./bin/redirect_data.json"
	}

	//Open file
	fileData, err := os.ReadFile(db.jsonFilepath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			// 1. Ensure the directory (e.g., "./bin") actually exists
			dir := filepath.Dir(db.jsonFilepath)
			if mkErr := os.MkdirAll(dir, 0755); mkErr != nil {
				return mkErr
			}

			// Initialize with an empty JSON structure "{}"
			fileData = []byte(`{"redirects":{}}`)

			if writeErr := os.WriteFile(db.jsonFilepath, fileData, 0644); writeErr != nil {
				return writeErr
			}

		} else {
			// Return any other unexpected error (like permission denied)
			return err
		}
	}

	//Export data
	//Save on db
	err = json.Unmarshal(fileData, &db)
	if err != nil {
		return err
	}
	return nil
}

func (db *JSONDB) Save() error {
	db.mu.Lock()
	defer db.mu.Unlock()
	//Prepare data
	//jsonData, err := json.MarshalIndent(db, "", "	")
	jsonData, err := json.Marshal(db)
	if err != nil {
		return err
	}
	//Open file
	err = os.WriteFile(db.jsonFilepath, jsonData, 0644)
	if err != nil {
		return err
	}
	//Save to file
	return nil
}

func (db *JSONDB) AddRedirect(newRedirect types.Redirect) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if !utils.IsRedirectValid(newRedirect) {
		return errors.New(enums.ERRORBadRedirectData)
	}

	_, inUse := db.Redirects[newRedirect.Slug]
	logger.Log(inUse)

	if inUse {
		return errors.New(enums.ERRORExistingSlug)
	}

	logger.Log(newRedirect)
	db.Redirects[newRedirect.Slug] = newRedirect

	logger.Log(db.Redirects)

	return nil

}
func (db *JSONDB) RemoveRedirect(slug types.Slug) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exists := db.Redirects[types.Slug(slug)]

	if !exists {
		return errors.New(enums.ERRORNoMatchingSlug)
	}

	delete(db.Redirects, slug)

	return nil
}
func (db *JSONDB) GetRedirects() ([]types.Redirect, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	returnRedirects := types.RedirectSlice(utils.RedirectMaptoSlice(db.Redirects))
	returnRedirects.Sort()

	return returnRedirects, nil
}
