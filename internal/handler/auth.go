package handler

import (
	"log/slog"
	"net/http"

	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/Tinddd28/GoPTL/pkg/logger"
	"github.com/Tinddd28/GoPTL/pkg/random"
	"github.com/Tinddd28/GoPTL/pkg/sender"
	"github.com/gin-gonic/gin"
)

// @Summary Register
// @Tags auth
// @Description register new models
// @ID create-account
// @Accept json
// @Produce json
// @Param input body models.User true "models info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid input")
	}

	input.Password = random.RandomPass(10)
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	log_ := logger.SetupPrettyLogger()
	log_.Info("User created", slog.Any("data", input))
	err = sender.SendMail(sender.Sender{
		Email:    input.Email,
		Pass:     input.Password,
		Name:     input.Name,
		LastName: input.Lastname,
	})
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to send email")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Login
// @Tags auth
// @Description login models
// @ID login
// @Accept json
// @Produce json
// @Param input body Login true "email and password"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var input Login
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Invalid input")
	}
	//log_ := logger.SetupPrettyLogger()
	//log_.Info("mail and pass: ", slog.Any("data", input))

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	//log_.Info(token)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
