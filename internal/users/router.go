package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/auth"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/permissions"
)

func SetupUserRoutes(g *gin.Engine, dbPool *pgxpool.Pool) {
	var (
		authRepository auth.IAuthRepository = auth.NewAuthRepository()
		userRepository IUserRepository      = NewUserRepository(dbPool, authRepository)
	)

	routerUsers := g.Group("/users")
	routerUsers.POST(
		"/",
		auth.OnlyAuthenticated(authRepository),
		permissions.OnlyPermission(permissions.PermissionCreateUser),
		handleCreateUser(userRepository),
	)
	routerUsers.GET(
		"/",
		auth.OnlyAuthenticated(authRepository),
		permissions.OnlyPermission(permissions.PermissionSeeAllUsers),
		handleGetUsers(userRepository),
	)

	routerAuth := g.Group("/auth")
	routerAuth.POST(
		"/",
		handleLogin(userRepository, authRepository),
	)
}
