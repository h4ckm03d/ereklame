// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: documents.sql

package sqlc

import (
	"context"
)

const deleteDocument = `-- name: DeleteDocument :exec
DELETE FROM documents WHERE id = $1
`

func (q *Queries) DeleteDocument(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteDocument, id)
	return err
}

const getDocument = `-- name: GetDocument :one
SELECT id, permit_id, document_url, created_at FROM documents WHERE id = $1
`

func (q *Queries) GetDocument(ctx context.Context, id int32) (Document, error) {
	row := q.db.QueryRowContext(ctx, getDocument, id)
	var i Document
	err := row.Scan(
		&i.ID,
		&i.PermitID,
		&i.DocumentUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getDocuments = `-- name: GetDocuments :many
SELECT id, permit_id, document_url, created_at FROM documents
`

func (q *Queries) GetDocuments(ctx context.Context) ([]Document, error) {
	rows, err := q.db.QueryContext(ctx, getDocuments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Document
	for rows.Next() {
		var i Document
		if err := rows.Scan(
			&i.ID,
			&i.PermitID,
			&i.DocumentUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const newDocument = `-- name: NewDocument :one
INSERT INTO documents (permit_id, document_url) VALUES ($1, $2) RETURNING id, permit_id, document_url, created_at
`

type NewDocumentParams struct {
	PermitID    int32  `json:"permit_id"`
	DocumentUrl string `json:"document_url"`
}

func (q *Queries) NewDocument(ctx context.Context, arg NewDocumentParams) (Document, error) {
	row := q.db.QueryRowContext(ctx, newDocument, arg.PermitID, arg.DocumentUrl)
	var i Document
	err := row.Scan(
		&i.ID,
		&i.PermitID,
		&i.DocumentUrl,
		&i.CreatedAt,
	)
	return i, err
}

const updateDocument = `-- name: UpdateDocument :one
UPDATE documents SET document_url = $1 WHERE id = $2 RETURNING id, permit_id, document_url, created_at
`

type UpdateDocumentParams struct {
	DocumentUrl string `json:"document_url"`
	ID          int32  `json:"id"`
}

func (q *Queries) UpdateDocument(ctx context.Context, arg UpdateDocumentParams) (Document, error) {
	row := q.db.QueryRowContext(ctx, updateDocument, arg.DocumentUrl, arg.ID)
	var i Document
	err := row.Scan(
		&i.ID,
		&i.PermitID,
		&i.DocumentUrl,
		&i.CreatedAt,
	)
	return i, err
}
