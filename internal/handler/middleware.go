package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
	superUserCtx        = "is_superuser"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, issuperuser, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid token")
	}

	c.Set(userIdCtx, userId)
	c.Set(superUserCtx, issuperuser)
}
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userIdCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func getSuperUser(c *gin.Context) (bool, error) {
	id, ok := c.Get(superUserCtx)
	if !ok {
		return false, errors.New("user id not found")
	}

	idBool, ok := id.(bool)
	if !ok {
		return false, errors.New("user id is of invalid type")
	}

	return idBool, nil
}
