package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/api-gerenciador-fila/config"
	database "github.com/mateusgcoelho/api-gerenciador-fila/database/sqlc"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/utils"
)

func HandleLogin(r *database.DatabaseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(fmt.Sprintf("Corpo da requisição invalido. %v", err), nil))
			return
		}

		user, err := r.GetUserByEmail(context.Background(), req.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(err.Error(), nil))
			return
		}

		if !ComparePasswords(user.Password, req.Password) {
			ctx.JSON(http.StatusUnauthorized, utils.BuildResponse("E-mail e/ou senha inválidos.", nil))
			return
		}

		token, err := generateJwtToken(JwtPayload{
			Id:          int(user.ID),
			Permissions: int(user.Permissions),
		}, config.AppConfig.JwtSecretKey)
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, utils.BuildResponse("Não foi possível realizar autenticação.", nil))
			return
		}

		ctx.JSON(http.StatusOK, utils.BuildResponse("", token))
	}
}
