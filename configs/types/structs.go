package types

import (
	"html/template"
)

type Redirect struct {
	Description string
	Slug        string
	Destination template.URL
}

type AdminData struct {
	Redirects []Redirect
}

type LogInInfo struct {
	Username string
	Password string
}
