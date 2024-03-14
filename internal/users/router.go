package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/permissions"
)

func SetupUserRoutes(g *gin.Engine, dbPool *pgxpool.Pool) {
	var (
		userRepository IUserRepository = NewUserRepository(dbPool)
	)

	r := g.Group("/users")

	r.POST(
		"/",
		permissions.OnlyPermission(permissions.PermissionCreateUser),
		handleCreateUser(userRepository),
	)
}
