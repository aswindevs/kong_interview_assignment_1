package utils

import (
	"github.com/aswindevs/kong_interview-assignment_1/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SetCookie(c *gin.Context, token string, cfg *config.Auth) {
	c.SetCookie(
		cfg.CookieName,
		token,
		cfg.TokenExpirationTime*60*60, // Convert hours to seconds
		"/",
		cfg.Domain,
		false, // Secure flag
		true,  // HttpOnly flag
	)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
