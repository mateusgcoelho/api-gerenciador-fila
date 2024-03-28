-- name: GetReports :many
SELECT * FROM reports
LIMIT $1 OFFSET $2;

-- name: CreateReport :one
INSERT INTO reports (number, person_id, responsive_id, queue_id)
VALUES ($1, $2, $3, $4) RETURNING *;


-- name: GetReportWithoutAFinish :one
SELECT * FROM reports
WHERE finish_at IS NULL AND (person_id = $1 OR responsive_id = $2) LIMIT 1;
