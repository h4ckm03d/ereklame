-- name: GetFees :many
SELECT * FROM fees;

-- name: GetFee :one
SELECT * FROM fees WHERE id = $1;

-- name: NewFee :one
INSERT INTO fees (permit_id, amount) VALUES ($1, $2) RETURNING *;

-- name: UpdateFee :one
UPDATE fees SET amount = $1 WHERE id = $2 RETURNING *;

-- name: DeleteFee :exec
DELETE FROM fees WHERE id = $1;