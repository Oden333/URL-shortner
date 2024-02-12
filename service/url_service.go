package service

import (
	"URL-shortener/repository"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
)

type URL_service struct {
	repo repository.URLRepo
}

func NewURLService(repo repository.URLRepo) *URL_service {
	return &URL_service{repo: repo}
}

func (s *URL_service) SaveURL(urlToSave string) (string, error) {
	/*
		//Проводим валидацию данных
		if err := validator.New().Struct(url); err != nil {
			logrus.Debugf("Validation failed")
			return "", fmt.Errorf("Failed to validate data")
		}
	*/

	//Создаём псевдоним для сохренения в БД
	//var id int64
	var alias string

	//Анонимная функция сокращения URL
	var generateShortURL func(url string) string
	generateShortURL = func(url string) string {
		hash := sha256.Sum256([]byte(url))
		encoded := base64.URLEncoding.EncodeToString(hash[:8])
		return encoded[:8]
	}
	alias = generateShortURL(urlToSave)

	//Кодируем URL для ДБ
	encodedURL := url.QueryEscape(urlToSave)
	_, dbAlias, err := s.repo.SaveURL(encodedURL, alias)
	return dbAlias, err
}

func (s *URL_service) GetByAlias(alias string) (string, error) {
	if alias == "" {
		return "", ErrorEmptyAlias
	}
	/*
		//Проводим валидацию данных
		if err := validator.New().Struct(alias); err != nil {
			logrus.Debugln("Validation failed with error: ", err)
			return "", fmt.Errorf("Failed to validate data")
		}
	*/

	return s.repo.GetByAlias(alias)
}
