-- name: GetPayments :many
SELECT * FROM payments;

-- name: GetPayment :one
SELECT * FROM payments WHERE id = $1;

-- name: NewPayment :one
INSERT INTO payments (user_id, permit_id, amount, status, payment_method) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdatePayment :one
UPDATE payments SET amount = $1, status = $2, payment_method = $3 WHERE id = $4 RETURNING *;

-- name: DeletePayment :exec
DELETE FROM payments WHERE id = $1;