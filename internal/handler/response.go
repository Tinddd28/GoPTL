package handler

import (
	"github.com/Tinddd28/GoPTL/pkg/logger"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	var log_ *slog.Logger
	log_ = logger.SetupPrettyLogger()
	log_.Error(message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}
