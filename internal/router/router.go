package router

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	users.SetupUserRoutes(r, dbPool)
	queues.SetupQueuesRoutes(r, dbPool)
	reports.SetupReportsRoutes(r, dbPool)

	if err := r.Run(serverPort); err != nil {
		log.Fatal("Não foi possível iniciar o serviço.")
	}
}
