package permissions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/utils"
)

func OnlyPermission(permissions ...Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		return
		userPermissions := Permission(c.GetInt("user_permissions"))

		hasPermission := false
		for _, permission := range permissions {
			if userPermissions&permission > 0 {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, utils.DefaultResponse{
				Message: "Sem permissão para executar essa ação.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
