package database

import (
	"errors"
	"fmt"
	"goslug/api/configs/types"
	"slices"
)

type MemoryStorage struct {
	redirects []types.Redirect
}

func (s *MemoryStorage) Load() error {
	// Sample Data
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
	return nil
}

func (s *MemoryStorage) AddRedirect(newRedirect types.Redirect) error {
	if newRedirect.Description != "" && newRedirect.Slug != "" && newRedirect.Destination != "" {
		fmt.Print("\tValid Redirect format\t")

		if isSlugUnique(newRedirect.Slug, s.redirects) {
			fmt.Print("\tUnique Slug\t")

			// add it to the storage. ATP using just local mememory
			s.redirects = append(s.redirects, newRedirect)

			fmt.Println("Redirect Saved")
			return nil

		} else {
			// Slug already exist
			fmt.Println("\tREQ:ERR SLUG ALREADY EXIST\t")
			return errors.New(types.ERRORExistingSlug)
		}

	} else {
		// Bad Redirect Data
		fmt.Println("\tREQ:ERR BAD REDIRECT\t")
		return errors.New(types.ERRORBadRedirectData)
	}
}

func (s *MemoryStorage) RemoveRedirect(slug string) error {
	for i, v := range s.redirects {
		if slug == v.Slug {
			s.redirects = slices.Delete(s.redirects, i, i+1)

			fmt.Println("\tDELETED")
			return nil
		}
	}
	return errors.New(types.ERRORNoMatchingSlug)
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
