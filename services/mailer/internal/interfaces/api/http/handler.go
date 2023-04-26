package http

import (
	"net/http"

	"github.com/avalonprod/eliteeld/mailer/internal/config"
	"github.com/avalonprod/eliteeld/mailer/internal/domain/services"
	"github.com/avalonprod/eliteeld/mailer/internal/interfaces/api/http/v1"
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

func (h *Handlers) Init(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	h.initAPI(r)
	return r
}

func (h *Handlers) initAPI(r *gin.Engine) {
	v1 := v1.NewHandlers(h.services)
	api := r.Group("/api")
	{
		v1.Init(api)
	}
}
