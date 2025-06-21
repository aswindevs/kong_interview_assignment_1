package middlewares

import (
	"time"

	"github.com/aswindevs/kong_interview-assignment_1/pkg/logger"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(log logger.Interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		log.Info(
			c.Request.Context(),
			"request handled",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", duration,
		)
	}
}
