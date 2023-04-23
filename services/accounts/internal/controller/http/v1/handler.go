package v1

import (
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

func (h *Handler) InitRoutes(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initRoutesUser(v1)
	}
}
