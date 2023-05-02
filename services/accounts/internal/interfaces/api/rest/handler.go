package rest

import (
	"net/http"

	"github.com/avalonprod/eliteeld/accounts/internal/config"
	"github.com/avalonprod/eliteeld/accounts/internal/core/services"
	v1 "github.com/avalonprod/eliteeld/accounts/internal/interfaces/api/rest/v1"
	"github.com/avalonprod/eliteeld/accounts/pkg/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     *services.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *services.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}

}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	// TODO SET MODE FROM CONFIG
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	v1 := v1.NewHandler(h.services, h.tokenManager)
	api := router.Group("/api")
	{
		v1.Init(api)
	}
}
