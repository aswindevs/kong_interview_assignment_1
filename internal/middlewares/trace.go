package middlewares

import (
	"context"

	"github.com/aswindevs/kong_interview-assignment_1/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		ctx := context.WithValue(c.Request.Context(), logger.TraceIDKey, traceID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
