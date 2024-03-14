package permissions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OnlyPermission(permissions ...Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		userPermissions := Permission(c.GetInt("user_permissions"))

		hasPermission := false
		for _, permission := range permissions {
			if userPermissions&permission > 0 {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
