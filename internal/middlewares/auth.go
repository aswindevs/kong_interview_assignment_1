package middlewares

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/aswindevs/kong_interview-assignment_1/config"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// PermissionMap defines the permissions for each role
var PermissionMap = map[string][]string{
	"admin": {"create:all", "read:all", "update:all", "delete:all"},
	"user":  {"read:all"},
}

type routePermission struct {
	method     string
	permission string
}

var apiPermissions = map[string]routePermission{
	"/v1/services":               {method: "GET", permission: "read:all"},
	"/v1/services/:id":           {method: "GET", permission: "read:all"},
	"/v1/services/:id/versions":  {method: "GET", permission: "read:all"},
	"/v1/services/":              {method: "POST", permission: "create:all"},
	"/v1/services/:id/versions/": {method: "POST", permission: "create:all"},
}

// AuthMiddleware verifies JWT token and extracts userId
func AuthMiddleware(cfg *config.Auth, userUsecase usecase.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if shouldSkipAuth(c, cfg) {
			c.Next()
			return
		}

		token, err := validateToken(c, cfg)
		if err != nil {
			handleAuthError(c, err)
			return
		}

		if err := setUserContext(c, token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			return
		}

		if !checkPermissions(c, userUsecase) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			return
		}

		c.Next()
	}
}

func checkPermissions(c *gin.Context, userUsecase usecase.UserUsecase) bool {
	route := c.FullPath()
	method := c.Request.Method

	permissionInfo, ok := apiPermissions[route]
	if !ok || permissionInfo.method != method {
		return true // No specific permission required for this route/method
	}
	requiredPermission := permissionInfo.permission

	userID, exists := c.Get("user_id")
	if !exists {
		return false
	}

	user, err := userUsecase.GetUserById(c, userID.(int))
	if err != nil {
		return false // Cannot verify permissions if user is not found
	}

	for _, role := range user.Roles {
		if permissions, ok := PermissionMap[role.Name]; ok {
			for _, p := range permissions {
				if p == requiredPermission {
					return true
				}
			}
		}
	}

	return false
}

func shouldSkipAuth(c *gin.Context, cfg *config.Auth) bool {
	if slices.Contains(cfg.ExcludePaths, c.Request.URL.Path) {
		c.Next()
		return true
	}
	return false
}

func validateToken(c *gin.Context, cfg *config.Auth) (*jwt.Token, error) {
	cookie, err := c.Cookie(cfg.CookieName)
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func setUserContext(c *gin.Context, token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid claims")
	}

	userID := int(claims["user_id"].(float64))
	c.Set("user_id", userID)
	return nil
}

func handleAuthError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
}
