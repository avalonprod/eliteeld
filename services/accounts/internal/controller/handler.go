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

func (h *Handler) CorsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	// TODO SET MODE FROM CONFIG
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		h.CorsMiddleware,
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
