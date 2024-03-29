package v1

import (
	"github.com/avalonprod/eliteeld/mailer/pkg/logger"
	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logger := logger.NewLogger()
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
