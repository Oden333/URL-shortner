package handler

import (
	"URL-shortener/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//Конфигурация страницы ответа
	router := gin.Default()

	//Конфигурация рутов
	url := router.Group("/")
	{
		//Метод для добавления записей
		url.POST("a/", h.createUrl)

		//Метод для получения записей
		url.GET("s/:alias", h.getUrl)
	}
	return router
}
