package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/utils"
)

func OnlyAuthenticated(authRepository IAuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		bearerToken := strings.Split(token, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, utils.DefaultResponse{
				Message: "Token de autenticação inválido.",
			})
			c.Abort()
			return
		}

		payload, err := authRepository.ValidateToken(bearerToken[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.DefaultResponse{
				Message: "Não foi possível validar token de autenticação.",
			})
			c.Abort()
			return
		}

		c.Set("user_id", payload.Id)
		c.Set("user_permissions", payload.Permissoes)

		c.Next()
	}
}
