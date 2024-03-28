package internal

import (
	"github.com/gin-gonic/gin"
	database "github.com/mateusgcoelho/api-gerenciador-fila/database/sqlc"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/queues"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/users"
)

type Server struct {
	DatabaseRepository *database.DatabaseRepository
}

func (s *Server) Run(port string) {
	router := gin.Default()

	router.GET("/users", users.HandleGetUsers(s.DatabaseRepository))
	router.POST("/users", users.HandleCreateUser(s.DatabaseRepository))

	router.GET("/queues", queues.HandleGetQueues(s.DatabaseRepository))
	router.POST("/queues", queues.HandleCreateQueue(s.DatabaseRepository))

	router.Run(port)
}
