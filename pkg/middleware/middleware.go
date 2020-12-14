// Package middleware implements middlewares for routers
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/COOLizh/task_repo/pkg/auth"
)

// AuthMiddleware checks if a valid token exists
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		id, err := auth.ParseToken(strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid token")
			return
		}
		c.Set("user_id", id)
		c.Next()
	}
}
