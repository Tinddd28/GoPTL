package handler

import (
	"log/slog"

	"github.com/Tinddd28/GoPTL/pkg/logger"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	var log_ *slog.Logger = logger.SetupPrettyLogger()
	log_.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
