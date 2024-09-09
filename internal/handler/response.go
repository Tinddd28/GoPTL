package handler

import (
	"github.com/Tinddd28/GoPTL/pkg/logger"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	var log_ *slog.Logger
	log_ = logger.SetupPrettyLogger()
	log_.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
