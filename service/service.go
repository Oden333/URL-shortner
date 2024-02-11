package service

import "URL-shortener/repository"

type URLService interface {
}

type Service struct {
	URLService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		URLService: NewURLService(repos.URLRepo),
	}
}
