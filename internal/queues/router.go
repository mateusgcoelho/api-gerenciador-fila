package queues

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/auth"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/permissions"
)

func SetupQueuesRoutes(g *gin.Engine, dbPool *pgxpool.Pool) {
	var (
		authDao  auth.IAuthDao = auth.NewAuthDao()
		queueDao IQueueDao     = NewQueueDao(dbPool)
	)

	r := g.Group("/queues", auth.OnlyAuthenticated(authDao))

	r.POST(
		"/",
		permissions.OnlyPermission(permissions.PermissionCreateQueue),
		handleCreateQueue(queueDao),
	)
	r.GET(
		"/",
		permissions.OnlyPermission(permissions.PermissionSeeAllQueues),
		handleGetQueues(queueDao),
	)
	r.GET(
		"/:id",
		permissions.OnlyPermission(permissions.PermissionSeeAllQueues),
		handleGetQueueById(queueDao),
	)
}
