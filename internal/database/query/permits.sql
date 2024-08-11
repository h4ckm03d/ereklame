-- name: GetPermits :many
SELECT * FROM permits;

-- name: GetPermit :one
SELECT * FROM permits WHERE id = $1;

-- name: NewPermit :one
INSERT INTO permits (user_id, description, status) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdatePermit :one
UPDATE permits SET description = $1, status = $2 WHERE id = $3 RETURNING *;

-- name: DeletePermit :exec
DELETE FROM permits WHERE id = $1;