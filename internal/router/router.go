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

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Initialize(dbPool *pgxpool.Pool) {
	r := gin.Default()
	r.Use(corsMiddleware())

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	users.SetupUserRoutes(r, dbPool)
	queues.SetupQueuesRoutes(r, dbPool)
	reports.SetupReportsRoutes(r, dbPool)

	if err := r.Run(serverPort); err != nil {
		log.Fatal("Não foi possível iniciar o serviço.")
	}
}
