package reports

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/auth"
)

func SetupReportsRoutes(g *gin.Engine, dbPool *pgxpool.Pool) {
	var (
		authDao   auth.IAuthDao = auth.NewAuthDao()
		reportDao IReportDao    = NewReportDao(dbPool)
	)

	r := g.Group("/reports", auth.OnlyAuthenticated(authDao))

	r.POST("/", handleCreateReport(reportDao))
}
