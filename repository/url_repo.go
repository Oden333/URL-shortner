package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type URL_repo struct {
	db *sqlx.DB
}

func NewURLRepo(db *sqlx.DB) *URL_repo {
	return &URL_repo{db: db}
}
func (r *URL_repo) SaveURL(urlToSave string, alias string) (int64, string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, "", err
	}

	var urlAlias string
	var urlId int64

	//Добавляем запись в БД( в случае уникального alias)
	createURLQuery := fmt.Sprintf("INSERT INTO %s (url,alias) values ($1, $2) ON CONFLICT (alias) DO NOTHING RETURNING id, alias", urlsTable)
	row := r.db.QueryRow(createURLQuery, urlToSave, alias)
	err = row.Scan(&urlId, &urlAlias)
	if err != nil {
		if err == sql.ErrNoRows {
			// Запись не была вставлена из-за ON CONFLICT DO NOTHING, продолжаем поиск существующей записи
			// Продолжаем выполнение кода
		} else {
			// Ошибка при выполнении запроса
			tx.Rollback()
			logrus.Debug("Error while inserting url: ", err)
			return 0, "", err
		}
	}

	//В случае дубликата находим нужную строчку
	if urlId == 0 {
		urlAlias = alias
		query := fmt.Sprintf(`SELECT id FROM %s WHERE alias = $1`, urlsTable)
		err := r.db.Get(&urlId, query, urlAlias)
		if err != nil {
			logrus.Debugf("Error while selecting URL by existing alias(%s), returning: %s", alias, err)
			return urlId, urlAlias, nil
		}
		logrus.Info("Got request with existing URL. Found data:", "Id: ", urlId, "", "Alias:", urlAlias, "")
	} else {
		logrus.Info("Successful saving URL: ", "", "Id: ", urlId, "Alias: ", urlAlias)
	}
	return urlId, urlAlias, nil
}

func (r *URL_repo) GetByAlias(alias string) (string, error) {

	var url string
	query := fmt.Sprintf(`SELECT url FROM %s WHERE alias = $1`, urlsTable)
	err := r.db.Get(&url, query, alias)
	if err != nil {
		logrus.Debugf("Error while selecting url by alias, returning: %s", err)
		return url, errors.New(fmt.Sprintf(" No urls with such Alias (%s)", alias))
	}

	return url, err

}
