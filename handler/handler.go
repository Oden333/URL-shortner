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
	auth := router.Group("/people")
	{

		auth.POST("/add")

	}

	return router
}
