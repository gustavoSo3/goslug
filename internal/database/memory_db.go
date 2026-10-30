package database

import (
	"errors"
	"goslug/configs/enums"
	"goslug/configs/types"
	"goslug/internal/utils"
	"sync"
)

type MemoryDB struct {
	redirects map[types.Slug]types.Redirect
	mu        sync.Mutex
}

func (db *MemoryDB) Load() error {
	db.mu.Lock()
	defer db.mu.Unlock()
	// Sample Data
	db.redirects = map[types.Slug]types.Redirect{
		"/lin": {
			Description: "To my linkedin",
			Slug:        "/lin",
			Destination: "https://www.linkedin.com/in/gustavoso3",
		},
		"/gh": {
			Description: "Github :)",
			Slug:        "/gh",
			Destination: "https://github.com/gustavoSo3",
		},
		"/cv": {
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
	db.mu.Lock()
	defer db.mu.Unlock()

	if utils.IsRedirectValid(newRedirect) {
		// Bad Redirect Data
		return errors.New(enums.ERRORBadRedirectData)
	}

	_, exist := db.redirects[newRedirect.Slug]
	if exist {
		// Slug already exist
		return errors.New(enums.ERRORExistingSlug)
	}

	db.redirects[newRedirect.Slug] = newRedirect
	return nil

}

func (db *MemoryDB) RemoveRedirect(slug types.Slug) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exist := db.redirects[slug]
	if !exist {
		// Slug not in db
		return errors.New(enums.ERRORNoMatchingSlug)

	}

	delete(db.redirects, slug)

	return nil
}

func (db *MemoryDB) GetRedirects() ([]types.Redirect, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	returnRedirects := types.RedirectSlice(utils.RedirectMaptoSlice(db.redirects))
	returnRedirects.Sort()

	return returnRedirects, nil
}
