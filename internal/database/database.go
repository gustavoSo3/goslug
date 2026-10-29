package database

import "goslug/configs/types"

type DataBase interface {
	Load() error
	Save() error
	AddRedirect(types.Redirect) error
	RemoveRedirect(types.Slug) error
	GetRedirects() ([]types.Redirect, error)
}

var RedirectDB DataBase
