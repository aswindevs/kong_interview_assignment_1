package controller

import (
	"net/http"

	appError "github.com/aswindevs/kong_interview-assignment_1/internal/errors"
	"github.com/aswindevs/kong_interview-assignment_1/pkg/logger"

	"github.com/gin-gonic/gin"
)

// HandleError is a utility function to handle AppError and respond appropriately
func HandleError(c *gin.Context, logger logger.Interface, err error) {
	logger.Error(c, "request_failed", "error", err)
	if appErr, ok := err.(*appError.AppError); ok {
		c.JSON(appErr.HTTPCode, gin.H{
			"error":   appErr.Code,
			"message": appErr.Message,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
