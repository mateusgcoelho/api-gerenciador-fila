package queues

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/auth"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/permissions"
)

func SetupQueuesRoutes(g *gin.Engine, dbPool *pgxpool.Pool) {
	var (
		authRepository  auth.IAuthRepository = auth.NewAuthRepository()
		queueRepository IQueueRepository     = NewQueueRepository(dbPool)
	)

	r := g.Group("/queues", auth.OnlyAuthenticated(authRepository))

	r.POST(
		"/",
		permissions.OnlyPermission(permissions.PermissionCreateQueue),
		handleCreateQueue(queueRepository),
	)
	r.GET(
		"/",
		permissions.OnlyPermission(permissions.PermissionSeeAllQueues),
		handleGetQueues(queueRepository),
	)
}
