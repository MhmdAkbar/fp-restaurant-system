package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Role(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, ok := c.Get("user_role")
		if !ok || userRole != role {
			c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}
