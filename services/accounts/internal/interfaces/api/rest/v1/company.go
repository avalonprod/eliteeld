package v1

import (
	"errors"
	"net/http"

	"github.com/avalonprod/eliteeld/accounts/internal/core/types"
	"github.com/gin-gonic/gin"
)

const (
	ErrCompanyAlreadyExists = "user with such email already exists"
	ErrCompanyNotFound      = "user doesn't exists"
)

func (h *Handler) InitCompanyRoutes(api *gin.RouterGroup) {
	company := api.Group("company")
	{
		company.POST("/sign-up", h.companySignUp)
		company.POST("/sign-in", h.companySignIn)
		company.POST("/refresh", h.companyRefresh)
		authenticated := company.Group("/", h.companyIdentity)
		{
			// Remove
			authenticated.GET("/test-auth", func(c *gin.Context) { c.String(http.StatusOK, "авторизирован") })
		}
	}
}

func (h *Handler) companySignUp(c *gin.Context) {
	var input types.CompanySignUpDTO
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	err := h.services.Company.CompanySignUp(c.Request.Context(), input)
	if err != nil {
		if errors.Is(err, errors.New(ErrCompanyAlreadyExists)) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) companySignIn(c *gin.Context) {
	var input types.CompanySignIpDTO
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	res, err := h.services.Company.CompanySignIn(c.Request.Context(), input)
	if err != nil {
		if errors.Is(err, errors.New(ErrCompanyNotFound)) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.JSON(http.StatusOK, types.Tokens{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

func (h *Handler) companyRefresh(c *gin.Context) {
	var input types.RefreshToken
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	res, err := h.services.Company.RefreshTokens(c.Request.Context(), input.Token)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, types.Tokens{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}
