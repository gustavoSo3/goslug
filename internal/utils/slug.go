package utils

import (
	"goslug/configs/types"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Performs checks to know if a redirect has a valid format and all fields
func IsRedirectValid(redirect types.Redirect) bool {

	// If any field is empty its not valid
	if redirect.Description == "" || redirect.Slug == "" || redirect.Destination == "" {
		return false
	}

	// TODO: better format

	// Description is valid if is less than 100 chars
	var validDescription bool = utf8.RuneCountInString(redirect.Description) <= 100

	// Slug is valid if starts with "/" is less than 20 chars and doesnt have any spaces (maybe will need to check for -/_() later)
	var validSlug bool = strings.HasPrefix(string(redirect.Slug), "/") && utf8.RuneCountInString(string(redirect.Slug)) <= 20 && !strings.Contains(string(redirect.Slug), " ")

	// Must parse without error AND have both a scheme and a host
	u, err := url.Parse(redirect.Destination)
	var validDestination bool = (err == nil && u.Scheme != "" && u.Host != "")

	return validDescription && validSlug && validDestination
}
