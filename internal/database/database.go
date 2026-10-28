package database

import "goslug/configs/types"

type DataBase interface {
	Load() error
	Save() error
	AddRedirect(types.Redirect) error
	RemoveRedirect(string) error
	GetRedirects() ([]types.Redirect, error)
}

var RedirectDB DataBase
