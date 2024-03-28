package queues

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/mateusgcoelho/api-gerenciador-fila/database/sqlc"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/utils"
)

func HandleCreateQueue(r *database.DatabaseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var createQueueRequest CreateQueueRequest
		if err := ctx.ShouldBindJSON(&createQueueRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(fmt.Sprintf("Corpo da requisição invalido. %v", err), nil))
			return
		}

		queue, err := r.CreateQueue(context.Background(), createQueueRequest.Name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildResponse(err.Error(), nil))
			return
		}

		ctx.JSON(http.StatusOK, utils.BuildResponse("", queue))
	}
}

func HandleGetQueues(r *database.DatabaseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetQueuesRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(fmt.Sprintf("Requisição invalida. %v", err), nil))
			return
		}

		arg := database.GetQueuesParams{
			Limit:  req.PageSize,
			Offset: (req.Page - 1) * req.PageSize,
		}
		queues, err := r.GetQueues(context.Background(), arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(err.Error(), nil))
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildResponse("", queues))
	}
}
