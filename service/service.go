package service

import "URL-shortener/repository"

type URLService interface {
	SaveURL(url string) (string, error)
	GetByAlias(alias string) (string, error)
}

type Service struct {
	URLService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		URLService: NewURLService(repos.URLRepo),
	}
}
