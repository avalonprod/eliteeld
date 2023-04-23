package v1

import (
	"net/http"

	"github.com/avalonprod/eliteeld/accounts/internal/domain/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initRoutesUser(api *gin.RouterGroup) {
	accounts := api.Group("/accounts")
	company := accounts.Group("/company")
	{
		company.POST("/register", h.UserRegister)
		login := company.Group("/login")
		{
			login.POST("/email", h.UserLoginEmail)
			login.POST("/password", h.UserLoginPassword)
		}
	}

}

type UserLoginEmailResponse struct {
	Email string `json:"email"`
}

func (h *Handler) UserRegister(c *gin.Context) {
	var input model.RegisterUserInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.service.User.UserRegister(c.Request.Context(), input); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) UserLoginEmail(c *gin.Context) {
	var input model.LoginEmailUserInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	email, err := h.service.User.UserLoginEmail(c.Request.Context(), input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.JSON(http.StatusOK, UserLoginEmailResponse{
		Email: email,
	})
}

func (h *Handler) UserLoginPassword(c *gin.Context) {
	var input model.LoginPasswordUserInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	payload, err := h.service.User.UserLoginPassword(c.Request.Context(), input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.JSON(http.StatusOK, payload)
}
