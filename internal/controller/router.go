package controller

import (
	"github.com/aswindevs/kong_interview-assignment_1/config"
	"github.com/aswindevs/kong_interview-assignment_1/internal/dependencies"
	"github.com/aswindevs/kong_interview-assignment_1/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AddRouter(svc *dependencies.Dependencies, config *config.Config) {
	server := svc.Server

	// Add logging middleware

	server.Use(middlewares.TraceMiddleware())
	server.GET("/health", GetHealth)
	server.Use(middlewares.LoggerMiddleware(svc.Logger))
	server.Use(middlewares.AuthMiddleware(&config.Auth, *svc.UserUseCase))
	apiGroup := server.Group("/v1")
	initializeRouters(apiGroup, svc)
}

func initializeRouters(server *gin.RouterGroup, svc *dependencies.Dependencies) {
	authRouter := server.Group("/auth")
	serviceRouter := server.Group("/services")
	NewAuthController(authRouter, svc.UserUseCase, svc.AuthUseCase, svc.Logger, &svc.Config.Auth)
	NewServiceCatalogController(serviceRouter, svc.ServiceCatalogUseCase, svc.Logger, &svc.Config.Auth)
}
