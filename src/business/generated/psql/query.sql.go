// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package psql

import (
	"context"
	"database/sql"
	"time"
)

const countUser = `-- name: CountUser :one
SELECT COUNT(id) as total FROM users
`

func (q *Queries) CountUser(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUser)
	var total int64
	err := row.Scan(&total)
	return total, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
	id,
	name,
	email,
	password,
	profile_image,
	created_at,
	created_by,
	is_deleted
) VALUES (
	 $1, $2, $3, $4, $5, $6, $7, $8
) 
RETURNING id, name, email, password, profile_image, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted
`

type CreateUserParams struct {
	ID           string
	Name         string
	Email        string
	Password     string
	ProfileImage string
	CreatedAt    time.Time
	CreatedBy    string
	IsDeleted    bool
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.ProfileImage,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.IsDeleted,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.ProfileImage,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.IsDeleted,
	)
	return i, err
}

const getListUserWithPagination = `-- name: GetListUserWithPagination :many
SELECT id, name, email, password, profile_image, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted FROM users LIMIT $1 OFFSET $2
`

type GetListUserWithPaginationParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetListUserWithPagination(ctx context.Context, arg GetListUserWithPaginationParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getListUserWithPagination, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.ProfileImage,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.DeletedAt,
			&i.DeletedBy,
			&i.IsDeleted,
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

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, profile_image, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.ProfileImage,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.IsDeleted,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, email, password, profile_image, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.ProfileImage,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.IsDeleted,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET
	name = $2,
	email = $3,
	password = $4,
	profile_image = $5,
	updated_at = $6,
	updated_by = $7
WHERE
	id = $1
RETURNING id, name, email, password, profile_image, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted
`

type UpdateUserParams struct {
	ID           string
	Name         string
	Email        string
	Password     string
	ProfileImage string
	UpdatedAt    sql.NullTime
	UpdatedBy    sql.NullString
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.ProfileImage,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.ProfileImage,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.IsDeleted,
	)
	return i, err
}
