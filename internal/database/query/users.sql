-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: NewUser :one
INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET email = $1, name = $2, password = $3 WHERE id = $4 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;