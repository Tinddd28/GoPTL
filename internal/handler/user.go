package handler

import (
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Update user
// @Security ApiKeyAuth
// @Tags user
// @Description update user info
// @ID update-user
// @Accept json
// @Produce json
// @Param input body models.User true "user info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /user [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var input models.User
	id, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Usr.UpdateUsr(id, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get user
// @Security ApiKeyAuth
// @Tags user
// @Description get user info
// @ID get-user
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /user [get]
func (h *Handler) GetUsr(c *gin.Context) {
	var usr models.User
	id, err := getUserId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if usr, err = h.services.Usr.GetUsr(id); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, usr)
}
