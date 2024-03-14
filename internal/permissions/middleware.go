package permissions

import (
	"github.com/gin-gonic/gin"
)

func OnlyPermission(permissions ...Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
