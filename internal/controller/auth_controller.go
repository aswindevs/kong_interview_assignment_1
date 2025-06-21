package controller

import (
	"net/http"

	"github.com/aswindevs/kong_interview-assignment_1/config"
	appError "github.com/aswindevs/kong_interview-assignment_1/internal/errors"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase"
	"github.com/aswindevs/kong_interview-assignment_1/internal/utils"
	"github.com/aswindevs/kong_interview-assignment_1/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUsecase *usecase.AuthUsecase
	userUsecase *usecase.UserUsecase
	logger      logger.Interface
	cfg         *config.Auth
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthController(router *gin.RouterGroup, userUsercase *usecase.UserUsecase,
	authUsecase *usecase.AuthUsecase, logger logger.Interface, cfg *config.Auth) {
	controller := &AuthController{
		authUsecase: authUsecase,
		userUsecase: userUsercase,
		logger:      logger,
		cfg:         cfg,
	}
	router.POST("/login", controller.Login)
}

func (a *AuthController) Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		HandleError(c, a.logger, appError.NewBadRequestError("invalid request body"))
		return
	}

	token, err := a.authUsecase.Login(c.Request.Context(), loginReq.Email, loginReq.Password)
	if err != nil {
		HandleError(c, a.logger, err)
		return
	}
	utils.SetCookie(c, token, a.cfg)

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}
