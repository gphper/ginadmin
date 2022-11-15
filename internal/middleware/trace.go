package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/**
生成RequestId
*/
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := context.WithValue(c.Request.Context(), "requestId", uuid.New().String())
		c.Set("ctx", ctx)

		c.Next()
	}
}
