package middleware

import (
	"net/http"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

// AdminMiddleware ensures the authenticated user has admin privileges.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userVal, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		user, ok := userVal.(schemas.User)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user in context"})
			return
		}

		if !user.IsAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin privileges required"})
			return
		}
		
		c.Next()
	}
}