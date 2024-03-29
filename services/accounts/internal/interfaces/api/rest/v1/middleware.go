package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"

	companyCtx = "companyId"
)

func (h *Handler) companyIdentity(c *gin.Context) {
	id, err := h.parseAuthHeader(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(companyCtx, id)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}

func getCompanyId(c *gin.Context) (string, error) {
	return getIdByContext(c, companyCtx)
}

func getIdByContext(c *gin.Context, context string) (string, error) {
	idFromCtx, ok := c.Get(context)
	if !ok {
		return "", errors.New("companyCtx not found")
	}
	id := idFromCtx.(string)
	return id, nil
}
