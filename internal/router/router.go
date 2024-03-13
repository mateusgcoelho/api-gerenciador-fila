package router

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	if err := r.Run(serverPort); err != nil {
		log.Fatal("Não foi possível iniciar o serviço.")
	}
}
