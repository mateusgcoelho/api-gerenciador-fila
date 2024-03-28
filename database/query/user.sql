-- name: CreateUser :one
INSERT INTO users (email, password, code, permissions, person_id)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUsers :many
SELECT * FROM users
LIMIT $1 OFFSET $2;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;
