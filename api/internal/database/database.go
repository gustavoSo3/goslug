package database

import "goslug/api/configs/types"

type DataBase interface {
	Load() error
	Save([]types.Redirect) error
	AddRedirect(types.Redirect) error
	RemoveRedirect(string) error
	GetRedirects() ([]types.Redirect, error)
}

var RedirectDB DataBase
