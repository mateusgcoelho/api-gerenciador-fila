package database

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func RunMigrations() error {
	migrationFile, err := os.ReadFile("migrations.sql")
	if err != nil {
		return err
	}

	queries := strings.Split(string(migrationFile), ";")

	for _, query := range queries {
		_, err := dbPool.Exec(context.Background(), query)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}
