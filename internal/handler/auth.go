package handler

import (
	"github.com/Tinddd28/GoPTL/internal/user"
	"github.com/Tinddd28/GoPTL/pkg/logger"
	"github.com/Tinddd28/GoPTL/pkg/random"
	"github.com/Tinddd28/GoPTL/pkg/sender"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) Register(c *gin.Context) {
	var input user.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	input.Password = random.RandomPass(10)
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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

func (h *Handler) Login(c *gin.Context) {
	var input Login
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	//log_ := logger.SetupPrettyLogger()
	//log_.Info("mail and pass: ", slog.Any("data", input))

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	//log_.Info(token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
