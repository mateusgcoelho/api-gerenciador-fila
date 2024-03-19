package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/auth"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/permissions"
)

func SetupUserRoutes(g *gin.Engine, dbPool *pgxpool.Pool) {
	var (
		authDao auth.IAuthDao = auth.NewAuthDao()
		userDao IUserDao      = NewUserDao(dbPool, authDao)
	)

	routerUsers := g.Group("/users")
	routerUsers.POST(
		"/",
		permissions.OnlyPermission(permissions.PermissionCreateUser),
		handleCreateUser(userDao),
	)
	routerUsers.GET(
		"/",
		permissions.OnlyPermission(permissions.PermissionSeeAllUsers),
		handleGetUsers(userDao),
	)

	routerAuth := g.Group("/auth")
	routerAuth.POST(
		"/",
		handleLogin(userDao, authDao),
	)
}
