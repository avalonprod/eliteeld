package v1

import (
	"fmt"
	"net/http"

	"github.com/avalonprod/eliteeld/mailer/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) initEmails(api *gin.RouterGroup) {
	email := api.Group("/email")
	{
		email.POST("/company-registration", h.SendEmailCompanyRegistration)
	}
}

func (h *Handlers) SendEmailCompanyRegistration(c *gin.Context) {
	var input models.CompanyRegistrationEmailDTO

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Emails.SendEmailCompanyRegistration(c.Request.Context(), input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to send message from email: %s", input.Email))
		return
	}
	c.JSON(http.StatusOK, response{"success"})
}
