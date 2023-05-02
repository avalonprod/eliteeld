package v1

import (
	"github.com/avalonprod/eliteeld/accounts/internal/core/services"
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

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.InitCompanyRoutes(v1)
	}
}
