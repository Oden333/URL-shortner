package repository

import (
	"github.com/jmoiron/sqlx"
)

type URL_repo struct {
	db *sqlx.DB
}

func NewURLRepo(db *sqlx.DB) *URL_repo {
	return &URL_repo{db: db}
}
