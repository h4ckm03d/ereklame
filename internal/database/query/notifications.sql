-- name: GetNotifications :many
SELECT * FROM notifications;

-- name: GetNotification :one
SELECT * FROM notifications WHERE id = $1;

-- name: NewNotification :one
INSERT INTO notifications (user_id, message, status) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateNotification :one
UPDATE notifications SET message = $1, status = $2 WHERE id = $3 RETURNING *;

-- name: DeleteNotification :exec
DELETE FROM notifications WHERE id = $1;