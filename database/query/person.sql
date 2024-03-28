-- name: CreatePerson :one
INSERT INTO persons (name, phone, cpf)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetPersons :many
SELECT * FROM persons
LIMIT $1 OFFSET $2;

-- name: GetPerson :one
SELECT * FROM persons WHERE id = $1 LIMIT 1;
