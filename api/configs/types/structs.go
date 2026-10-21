package types

import (
	"html/template"
)

// Redirect struct
// Use to store basic redirec data, like a small description and 2 urls. The "slug" that we had, and the Destination
type Redirect struct {
	Description string
	Slug        template.URL
	Destination template.URL
}

// AdminData struct
// AdminData struct stores the information that is needed to compile the admin.html template
type AdminData struct {
	Redirects []Redirect
	Host      string
}

// //Sample data

// TODO: NOTE  I might have to change slug to start with out the / root TBD
var SampleAdminData AdminData = AdminData{
	Redirects: []Redirect{
		{
			Description: "first_redirect",
			Slug:        "/linkedin",
			Destination: "https://www.linkedin.com/in/gustavoso3",
		},
		{
			Description: "Github :)",
			Slug:        "/gh",
			Destination: "https://github.com/gustavoSo3",
		},
		{
			Description: "togoogle",
			Slug:        "/ggl",
			Destination: "https://www.google.com",
		},
		{
			Description: "youtubet",
			Slug:        "/yt",
			Destination: "https://www.linkedin.com/in/gustavoso3",
		},
	},
	Host: "localhost:8080",
}
