package database

import (
	"errors"
	"goslug/configs/enums"
	"goslug/configs/types"
	"goslug/internal/utils"
	"slices"
)

type MemoryDB struct {
	redirects []types.Redirect
}

func (db *MemoryDB) Load() error {
	// Sample Data
	db.redirects = []types.Redirect{
		{
			Description: "To my linkedin",
			Slug:        "/lin",
			Destination: "https://www.linkedin.com/in/gustavoso3",
		},
		{
			Description: "Github :)",
			Slug:        "/gh",
			Destination: "https://github.com/gustavoSo3",
		},
		{
			Description: "To my CV",
			Slug:        "/cv",
			Destination: "https://storage.googleapis.com/resume_sotres_dev/gustavo_sotres_oct_2026.pdf",
		},
	}
	return nil
}

func (db *MemoryDB) Save() error {
	return nil
}

func (db *MemoryDB) AddRedirect(newRedirect types.Redirect) error {

	if utils.IsRedirectValid(newRedirect) {
		// Bad Redirect Data
		return errors.New(enums.ERRORBadRedirectData)
	}

	_, exist := utils.GetSlugIndex(newRedirect.Slug, db.redirects)
	if exist {
		// Slug already exist
		return errors.New(enums.ERRORExistingSlug)
	}

	db.redirects = append(db.redirects, newRedirect)

	return nil

}

func (db *MemoryDB) RemoveRedirect(slug string) error {

	index, exist := utils.GetSlugIndex(slug, db.redirects)
	if !exist {
		// Slug not in db
		return errors.New(enums.ERRORNoMatchingSlug)

	}

	db.redirects = slices.Delete(db.redirects, index, index+1)

	return nil
}

func (db *MemoryDB) GetRedirects() ([]types.Redirect, error) {
	return db.redirects, nil
}
