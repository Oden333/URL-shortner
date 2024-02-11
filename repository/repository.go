package repository

import (
	"github.com/jmoiron/sqlx"
)

type URLRepo interface {
}

type Repository struct {
	URLRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		URLRepo: NewURLRepo(db),
	}
}
