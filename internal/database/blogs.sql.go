// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: blogs.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createBlogs = `-- name: CreateBlogs :one
INSERT INTO blogs (id, created_at, updated_at, body, title, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, body, title, visibility, user_id
`

type CreateBlogsParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Body      string
	Title     string
	UserID    uuid.UUID
}

func (q *Queries) CreateBlogs(ctx context.Context, arg CreateBlogsParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, createBlogs,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Body,
		arg.Title,
		arg.UserID,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.Title,
		&i.Visibility,
		&i.UserID,
	)
	return i, err
}

const deleteBlog = `-- name: DeleteBlog :one
DELETE FROM blogs
WHERE id = $1
RETURNING id, created_at, updated_at, body, title, visibility, user_id
`

func (q *Queries) DeleteBlog(ctx context.Context, id uuid.UUID) (Blog, error) {
	row := q.db.QueryRowContext(ctx, deleteBlog, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.Title,
		&i.Visibility,
		&i.UserID,
	)
	return i, err
}

const getAllTypeBlogs = `-- name: GetAllTypeBlogs :one
SELECT id, created_at, updated_at, body, title, visibility, user_id FROM blogs WHERE visibility=$1
`

func (q *Queries) GetAllTypeBlogs(ctx context.Context, visibility uuid.UUID) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getAllTypeBlogs, visibility)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.Title,
		&i.Visibility,
		&i.UserID,
	)
	return i, err
}

const getBlog = `-- name: GetBlog :one
SELECT id, created_at, updated_at, body, title, visibility, user_id FROM blogs WHERE id=$1
`

func (q *Queries) GetBlog(ctx context.Context, id uuid.UUID) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlog, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.Title,
		&i.Visibility,
		&i.UserID,
	)
	return i, err
}

const getUserBlogs = `-- name: GetUserBlogs :one
SELECT id, created_at, updated_at, body, title, visibility, user_id FROM blogs WHERE user_id = $1
`

func (q *Queries) GetUserBlogs(ctx context.Context, userID uuid.UUID) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getUserBlogs, userID)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.Title,
		&i.Visibility,
		&i.UserID,
	)
	return i, err
}

const updateUserBlog = `-- name: UpdateUserBlog :one
UPDATE blogs
SET body = $1, title = $2, visibility = $3
WHERE id = $4
RETURNING id, created_at, updated_at, body, title, visibility, user_id
`

type UpdateUserBlogParams struct {
	Body       string
	Title      string
	Visibility uuid.UUID
	ID         uuid.UUID
}

func (q *Queries) UpdateUserBlog(ctx context.Context, arg UpdateUserBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, updateUserBlog,
		arg.Body,
		arg.Title,
		arg.Visibility,
		arg.ID,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.Title,
		&i.Visibility,
		&i.UserID,
	)
	return i, err
}
