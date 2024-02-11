package service

import "URL-shortener/repository"

type URL_service struct {
	repo repository.URLRepo
}

func NewURLService(repo repository.URLRepo) *URL_service {
	return &URL_service{repo: repo}
}
