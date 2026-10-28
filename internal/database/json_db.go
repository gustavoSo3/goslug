package database

import (
	"encoding/json"
	"errors"
	"goslug/configs/enums"
	"goslug/configs/types"
	"goslug/internal/utils"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

type JSONDB struct {
	Redirects    []types.Redirect `json:"redirects"`
	jsonFilepath string
}

func (db *JSONDB) Load() error {
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
			fileData = []byte("{}")

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

	if !utils.IsRedirectValid(newRedirect) {
		return errors.New(enums.ERRORBadRedirectData)
	}

	_, exists := utils.GetSlugIndex(newRedirect.Slug, db.Redirects)
	if exists {
		return errors.New(enums.ERRORExistingSlug)
	}

	db.Redirects = append(db.Redirects, newRedirect)
	db.Save()

	return nil

}
func (db *JSONDB) RemoveRedirect(slug string) error {
	i, exists := utils.GetSlugIndex(slug, db.Redirects)

	if !exists {
		return errors.New(enums.ERRORNoMatchingSlug)
	}

	db.Redirects = slices.Delete(db.Redirects, i, i+1)
	db.Save()

	return nil
}
func (db *JSONDB) GetRedirects() ([]types.Redirect, error) {
	return db.Redirects, nil
}
