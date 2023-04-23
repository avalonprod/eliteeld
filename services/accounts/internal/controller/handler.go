package controller

import (
	"net/http"

	"github.com/avalonprod/eliteeld/accounts/internal/config"
	v1 "github.com/avalonprod/eliteeld/accounts/internal/controller/http/v1"
	"github.com/avalonprod/eliteeld/accounts/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	// TODO SET MODE FROM CONFIG
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	v1 := v1.NewHandler(h.service)
	api := router.Group("/api")
	{
		v1.InitRoutes(api)
	}
}
