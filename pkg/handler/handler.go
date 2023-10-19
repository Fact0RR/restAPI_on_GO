package handler

import (
	"github.com/Ubludor/restAPI_on_GO/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up",h.singUp)
		auth.POST("/sing-in",h.singIn)
	}

	api := router.Group("/api")
	{
		api.GET("/",h.getTodayDuty)
		api.GET("/:day",h.getDutyByDay)
		api.POST("/trash",h.setTrash)
		api.POST("/dish",h.setDish)
		api.DELETE("/trash",h.deleteTrash)
		api.DELETE("/dish",h.deleteDish)
	}

	return router
}
