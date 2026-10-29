package utils

import "goslug/configs/types"

func RedirectMaptoSlice(redirects map[types.Slug]types.Redirect) []types.Redirect {

	var returnRedirects []types.Redirect

	for _, redirect := range redirects {
		returnRedirects = append(returnRedirects, redirect)
	}

	return returnRedirects
}
