package types

type Redirect struct {
	Description string
	Slug        string
	Destination string
}

type AdminData struct {
	Redirects []Redirect
}

type LogInInfo struct {
	Username string
	Password string
}
