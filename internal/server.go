package internal

import (
	"github.com/gin-gonic/gin"
	database "github.com/mateusgcoelho/api-gerenciador-fila/database/sqlc"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/auth"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/queues"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/reports"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/users"
)

type Server struct {
	DatabaseRepository *database.DatabaseRepository
}

func (s *Server) Run(port string) {
	router := gin.Default()

	router.POST("/auth", auth.HandleLogin(s.DatabaseRepository))

	router.GET("/users", users.HandleGetUsers(s.DatabaseRepository))
	router.POST("/users", users.HandleCreateUser(s.DatabaseRepository))

	router.GET("/queues", queues.HandleGetQueues(s.DatabaseRepository))
	router.POST("/queues", queues.HandleCreateQueue(s.DatabaseRepository))
	// Editar fila
	// Excluir fila

	router.GET("/reports", reports.HandleGetReports(s.DatabaseRepository))
	router.POST("/reports", reports.HandleCreateReport(s.DatabaseRepository))
	// Finalizar um atendimento
	// Editar atendimento
	// Excluir atendimento
	// Cancelar um atendimento?

	router.Run(port)
}
