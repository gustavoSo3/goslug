package database

import (
	"errors"
	"goslug/configs/enums"
	"goslug/configs/types"
	"goslug/internal/logger"
	"slices"
)

type MemoryStorage struct {
	redirects []types.Redirect
}

func (s *MemoryStorage) Load() error {
	// Sample Data
	logger.Log("MEMORY_STORAGE_LOAD", "Creating sample redirects")
	s.redirects = []types.Redirect{
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
			Description: "youtubet",
			Slug:        "/yt",
			Destination: "https://www.linkedin.com/in/gustavoso3",
		},
		{
			Description: "To my CV",
			Slug:        "/cv",
			Destination: "https://storage.googleapis.com/resume_sotres_dev/gustavo_sotres_oct_2026.pdf",
		},
	}
	return nil
}

func (s *MemoryStorage) Save(redirects []types.Redirect) error {
	logger.Log("MEMORY_STORAGE_SAVE", "Redirects saved")
	return nil
}

func (s *MemoryStorage) AddRedirect(newRedirect types.Redirect) error {
	logger.Log("MEMORY_STORAGE_NEW_REDIRECT", "Slug:", newRedirect.Slug)
	if newRedirect.Description != "" && newRedirect.Slug != "" && newRedirect.Destination != "" {
		logger.Log("MEMORY_STORAGE_NEW_REDIRECT", "Valid Format")

		if isSlugUnique(newRedirect.Slug, s.redirects) {
			logger.Log("MEMORY_STORAGE_NEW_REDIRECT", "Slug Available")

			// add it to the storage. ATP using just local mememory
			s.redirects = append(s.redirects, newRedirect)

			logger.Log("MEMORY_STORAGE_NEW_REDIRECT", "Saved")
			return nil

		} else {
			// Slug already exist
			logger.Error("MEMORY_STORAGE_NEW_REDIRECT", enums.ERRORExistingSlug)
			return errors.New(enums.ERRORExistingSlug)
		}

	} else {
		// Bad Redirect Data
		logger.Error("MEMORY_STORAGE_NEW_REDIRECT", enums.ERRORBadRedirectData)
		return errors.New(enums.ERRORBadRedirectData)
	}
}

func (s *MemoryStorage) RemoveRedirect(slug string) error {
	logger.Log("MEMORY_STORAGE_REMOVE_REDIRECT", "Slug:", slug)

	for i, v := range s.redirects {
		if slug == v.Slug {
			s.redirects = slices.Delete(s.redirects, i, i+1)

			logger.Log("MEMORY_STORAGE_REMOVE_REDIRECT", "Removed")
			return nil
		}
	}
	logger.Error("MEMORY_STORAGE_REMOVE_REDIRECT", enums.ERRORNoMatchingSlug)
	return errors.New(enums.ERRORNoMatchingSlug)
}

func (s *MemoryStorage) GetRedirects() ([]types.Redirect, error) {
	return s.redirects, nil
}

func isSlugUnique(slug string, redirects []types.Redirect) bool {
	for _, v := range redirects {
		if slug == v.Slug {
			return false
		}
	}
	return true
}
