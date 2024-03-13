package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbPool *pgxpool.Pool
)

func GetDbPool() *pgxpool.Pool {
	if dbPool == nil {
		log.Fatal("Instancia de banco de dados não foi encontrada.")
	}

	return dbPool
}

func InitializeDatabase() error {
	databaseUrl := getDatabaseUrl()

	newDbPool, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		return fmt.Errorf("Não foi possível realizar conexão com banco de dados: %v", err)
	}

	if err := newDbPool.Ping(context.Background()); err != nil {
		newDbPool.Close()
		return fmt.Errorf("Não foi possível realizar conexão com banco de dados: %v", err)
	}

	dbPool = newDbPool

	return nil
}

func getDatabaseUrl() string {
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, databaseName)
}
