package v1

import (
	"github.com/avalonprod/eliteeld/mailer/internal/domain/services"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		services: services,
	}
}

func (h *Handlers) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initEmails(v1)
	}
}
