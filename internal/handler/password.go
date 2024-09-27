package handler

import (
	"net/http"

	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/Tinddd28/GoPTL/pkg/random"
	"github.com/Tinddd28/GoPTL/pkg/sender"
	"github.com/gin-gonic/gin"
)

// @Summary Change password
// @Security ApiKeyAuth
// @Tags password
// @Description change password
// @ID change-password
// @Accept json
// @Produce json
// @Param input body models.Password true "password info"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /password/change [patch]
func (h *Handler) ChangePassword(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Unauthorized")
		return
	}

	var input models.Password
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err = h.services.Password.ChangePassword(id, input.OldPassword, input.NewPassword); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to change password")
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// @Summary Reset password
// @Tags password
// @Description reset password
// @ID reset-password
// @Accept json
// @Produce json
// @Param email body models.PassReset true "email"
// @Success 200 {string} string	"Password has been reset"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /password/reset [post]
func (h *Handler) ResetPassword(c *gin.Context) {
	var email models.PassReset
	if err := c.BindJSON(&email); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	pass := random.RandomPass(10)
	if err := h.services.Password.ResetPassword(pass, email.Email); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to reset password")
	}

	if err := sender.SendMesResPass(pass, email.Email); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to send email")
		return
	}

	c.JSON(http.StatusOK, "Password has been reset")
}
