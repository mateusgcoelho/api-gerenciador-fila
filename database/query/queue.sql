-- name: GetQueues :many
SELECT * FROM queues
LIMIT $1 OFFSET $2;

-- name: CreateQueue :one
INSERT INTO queues (name)
VALUES ($1) RETURNING *;

-- name: GetQueueById :one
SELECT * FROM queues
WHERE id = $1 LIMIT 1;
