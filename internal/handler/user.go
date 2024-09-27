package handler

import (
	"net/http"
	"strconv"

	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/Tinddd28/GoPTL/pkg/sender"
	"github.com/gin-gonic/gin"
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
		NewErrorResponse(c, http.StatusInternalServerError, "Unauthorized")
		return
	}
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	err = h.services.Usr.UpdateUsr(id, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to update user")
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
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /user/info [get]
func (h *Handler) GetUsr(c *gin.Context) {
	var usr models.UserResponse
	id, err := getUserId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if usr, err = h.services.Usr.GetUsr(id); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to get user")
		return
	}

	c.JSON(http.StatusOK, usr)
}

// @Summary Verification
// @Security ApiKeyAuth
// @Tags user
// @Description send verification code
// @ID verification
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /user/verification [post]
func (h *Handler) Verification(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var usr models.UserResponse

	if usr, err = h.services.Usr.GetUsr(id); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to get user")
		return
	}

	err = sender.SendVerification(id, usr.Email)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to send verification code")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    id,
		"email": usr.Email,
	})
}

// @Summary Apply verification
// @Tags user
// @Description apply verification code
// @ID apply-verification
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /user/verification_accept/{id} [get]
func (h *Handler) ApplyVerification(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Usr.Verification(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to verify user")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "verified",
	})
}
