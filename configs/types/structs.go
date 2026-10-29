package types

type Slug string

type Redirect struct {
	Description string
	Slug        Slug
	Destination string
}

type LogInInfo struct {
	Username string
	Password string
}
