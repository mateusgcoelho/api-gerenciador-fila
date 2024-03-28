package reports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	database "github.com/mateusgcoelho/api-gerenciador-fila/database/sqlc"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/utils"
)

func HandleCreateReport(r *database.DatabaseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateReportRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(fmt.Sprintf("Requisição invalida. %v", err), nil))
			return
		}

		person, err := r.GetPerson(context.Background(), int32(req.PersonID))
		if err != nil {
			if err == pgx.ErrNoRows {
				ctx.JSON(http.StatusBadRequest, utils.BuildResponse("Não foi possível encontrar pessoa.", nil))
				return
			}
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(err.Error(), nil))
			return
		}

		user, err := r.GetUserById(context.Background(), req.ResponsiveID)
		if err != nil {
			if err == pgx.ErrNoRows {
				ctx.JSON(http.StatusBadRequest, utils.BuildResponse("Não foi possível encontrar atendente.", nil))
				return
			}
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(err.Error(), nil))
			return
		}

		queue, err := r.GetQueueById(context.Background(), req.QueueID)
		if err != nil {
			if err == pgx.ErrNoRows {
				ctx.JSON(http.StatusBadRequest, utils.BuildResponse("Não foi possível encontrar fila.", nil))
				return
			}
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(err.Error(), nil))
			return
		}

		argGetReport := database.GetReportWithoutAFinishParams{
			PersonID: pgtype.Int4{
				Int32: person.ID,
				Valid: true,
			},
			ResponsiveID: user.ID,
		}
		reportExistent, err := r.GetReportWithoutAFinish(context.Background(), argGetReport)
		if err != nil && err != pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(err.Error(), nil))
			return
		}
		if reportExistent != (database.Report{}) {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse("Já existe um atendimento em andamento.", nil))
			return
		}

		arg := database.CreateReportParams{
			Number: 0,
			PersonID: pgtype.Int4{
				Int32: person.ID,
				Valid: true,
			},
			ResponsiveID: user.ID,
			QueueID:      queue.ID,
		}
		report, err := r.CreateReport(context.Background(), arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(err.Error(), nil))
			return
		}

		ctx.JSON(http.StatusOK, utils.BuildResponse("", report))
	}
}

func HandleGetReports(r *database.DatabaseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req GetReportsRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(fmt.Sprintf("Requisição invalida. %v", err), nil))
			return
		}

		arg := database.GetReportsParams{
			Limit:  req.PageSize,
			Offset: (req.Page - 1) * req.PageSize,
		}
		reports, err := r.GetReports(context.Background(), arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildResponse(err.Error(), nil))
			return
		}

		ctx.JSON(http.StatusOK, utils.BuildResponse("", reports))
	}
}
