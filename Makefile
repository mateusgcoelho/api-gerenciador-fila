run:
	go run ./cmd/api/main.go
migrate-up:
	migrate -path ./database/migrations/ -database postgres://postgres:Docker@localhost:5432?sslmode=disable up
migrate-down:
	migrate -path ./database/migrations/ -database postgres://postgres:Docker@localhost:5432?sslmode=disable down
.PHONY: run migrate-up migrate-down
