-- name: GetDocuments :many
SELECT * FROM documents;

-- name: GetDocument :one
SELECT * FROM documents WHERE id = $1;

-- name: NewDocument :one
INSERT INTO documents (permit_id, document_url) VALUES ($1, $2) RETURNING *;

-- name: UpdateDocument :one
UPDATE documents SET document_url = $1 WHERE id = $2 RETURNING *;

-- name: DeleteDocument :exec
DELETE FROM documents WHERE id = $1;