package router

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/queues"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/reports"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/users"
)

func Initialize(dbPool *pgxpool.Pool) {
	r := gin.Default()

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	users.SetupUserRoutes(r, dbPool)
	queues.SetupQueuesRoutes(r, dbPool)
	reports.SetupReportsRoutes(r, dbPool)

	if err := r.Run(serverPort); err != nil {
		log.Fatal("Não foi possível iniciar o serviço.")
	}
}
