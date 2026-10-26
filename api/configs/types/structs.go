package types

import (
	"html/template"
)

// Redirect struct
// Use to store basic redirec data, like a small description and 2 urls. The "slug" that we had, and the Destination
type Redirect struct {
	Description string
	Slug        string
	Destination template.URL
}

// AdminData struct
// AdminData struct stores the information that is needed to compile the admin.html template
type AdminData struct {
	Redirects []Redirect
}
