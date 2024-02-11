package repository

import (
	"github.com/jmoiron/sqlx"
)

type URLRepo interface {
	SaveURL(url, alias string) (int64, string, error)
	GetByAlias(alias string) (string, error)
}

type Repository struct {
	URLRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		URLRepo: NewURLRepo(db),
	}
}
