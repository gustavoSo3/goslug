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

// //Sample data

// NOTE: Slugs must be saved with the  "/" root slash. e.g. "/linkedin"
var SampleAdminData AdminData = AdminData{
	Redirects: []Redirect{
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
	},
}
