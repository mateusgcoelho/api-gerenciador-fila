package permissions

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func OnlyPermission(permissions ...Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(permissions)
		c.Next()
	}
}
