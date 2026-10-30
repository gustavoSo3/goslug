package types

type Slug string

type Redirect struct {
	Description string
	Slug        Slug
	Destination string
}

type LoginInfo struct {
	Username string
	Password string
}
