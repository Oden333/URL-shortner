package handler

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createUrl(c *gin.Context) {

	// Получаем объект запроса и анмаршаллим
	var req Request
	req.URL = c.Query("url")
	if req.URL == "" {
		logrus.Debug(ErrorEmptyURL)
		newErrorResponse(c, http.StatusInternalServerError, ErrorEmptyURL.Error())
		return
	}
	logrus.Info("Got URL shortening request:", req.URL)

	//Передаём в сервис и получаем ответ
	alias, err := h.services.URLService.SaveURL(req.URL)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"alias": alias,
	})
}

func (h *Handler) getUrl(c *gin.Context) {
	// Получаем объект запроса и анмаршаллим
	var req Request
	req.Alias = c.Param("alias")
	if req.Alias == "" {
		logrus.Debug(ErrorEmptyAlias)
		newErrorResponse(c, http.StatusInternalServerError, ErrorEmptyAlias.Error())
		return
	}
	logrus.Info("Got request with Alias: ", "", req.Alias)
	//Передаём в сервис и получаем ответ

	decodedURL, err := h.services.URLService.GetByAlias(req.Alias)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	req.URL, err = url.QueryUnescape(decodedURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	logrus.Infof("Got response from DB with allias %s. Redirect link: %s", req.Alias, req.URL)
	c.Redirect(http.StatusFound, req.URL)

}
