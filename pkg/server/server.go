package server

import (
	"github.com/aswindevs/kong_interview-assignment_1/pkg/logger"

	"github.com/gin-gonic/gin"
)

// New returns a configured gin engine
func New(log logger.Interface) *gin.Engine {
	// Disable gin's default logger
	gin.SetMode(gin.ReleaseMode)

	// Create a new engine without any middleware
	engine := gin.New()

	// Add recovery middleware
	engine.Use(gin.Recovery())

	return engine
}
