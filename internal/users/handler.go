package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	database "github.com/mateusgcoelho/api-gerenciador-fila/database/sqlc"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/utils"
)

func HandleCreateUser(r *database.DatabaseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(fmt.Sprintf("Corpo da requisição invalido. %v", err), nil))
			return
		}

		var user database.User

		err := r.ExecTx(context.Background(), func(q *database.Queries) error {
			argPerson := database.CreatePersonParams{
				Name: req.Name,
				Cpf: pgtype.Text{
					String: req.Cpf,
					Valid:  true,
				},
				Phone: pgtype.Text{
					String: req.Phone,
					Valid:  true,
				},
			}
			person, err := q.CreatePerson(context.Background(), argPerson)
			if err != nil {
				return err
			}

			argUser := database.CreateUserParams{
				Email:    req.Email,
				Password: req.Password,
				Code: pgtype.Text{
					String: req.Code,
					Valid:  true,
				},
				Permissions: 0,
				PersonID:    person.ID,
			}
			user, err = q.CreateUser(context.Background(), argUser)
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildResponse(err.Error(), nil))
			return
		}

		ctx.JSON(http.StatusOK, utils.BuildResponse("", user))
	}
}

func HandleGetUsers(r *database.DatabaseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetUsersRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(fmt.Sprintf("Requisição invalida. %v", err), nil))
			return
		}

		arg := database.GetUsersParams{
			Limit:  req.PageSize,
			Offset: (req.Page - 1) * req.PageSize,
		}
		users, err := r.GetUsers(context.Background(), arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(err.Error(), nil))
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildResponse("", users))
	}
}
