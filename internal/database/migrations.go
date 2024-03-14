package database

import (
	"context"
	"os"
	"strings"
)

func runMigrations() error {
	migrationFile, err := os.ReadFile("migrations.sql")
	if err != nil {
		return err
	}

	queries := strings.Split(string(migrationFile), ";")

	for _, query := range queries {
		_, err := dbPool.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}

	return nil
}
