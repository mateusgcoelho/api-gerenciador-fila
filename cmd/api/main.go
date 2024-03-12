package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar variaveis de ambiente.")
	}

	router.Initialize()
}
