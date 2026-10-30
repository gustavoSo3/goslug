package types

import (
	"cmp"
	"slices"
)

type Slug string

type Redirect struct {
	Description string
	Slug        Slug
	Destination string
}

type RedirectSlice []Redirect

// Sorts a redirect array in place
func (redirectSlice RedirectSlice) Sort() {
	slices.SortFunc(redirectSlice, func(a, b Redirect) int {
		return cmp.Compare(a.Slug, b.Slug)
	})
}

type LoginInfo struct {
	Username string
	Password string
}
