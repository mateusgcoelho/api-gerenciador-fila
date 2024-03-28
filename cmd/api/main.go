package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	db "github.com/mateusgcoelho/api-gerenciador-fila/database/sqlc"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal"
)

func main() {
	conn, err := pgxpool.New(context.Background(), "postgres://postgres:Docker@localhost:5432?sslmode=disable")
	if err != nil {
		panic(err)
	}
	dbRepository := db.NewDatabaseRepository(conn)

	server := internal.Server{
		DatabaseRepository: dbRepository,
	}

	server.Run(":3333")
}
