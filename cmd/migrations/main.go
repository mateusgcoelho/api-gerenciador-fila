package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar variaveis de ambiente.")
	}

	if err := database.InitializeDatabase(); err != nil {
		log.Fatal(err.Error())
	}

	dbPool := database.GetDbPool()

	database.RunMigrations()

	defer dbPool.Close()
}
