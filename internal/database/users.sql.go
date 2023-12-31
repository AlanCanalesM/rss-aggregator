// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, apikey, username, password) 
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'), $5, $6) RETURNING id, created_at, updated_at, name, apikey, username, password
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Username  string
	Password  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Username,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Apikey,
		&i.Username,
		&i.Password,
	)
	return i, err
}

const getUserByAPIKey = `-- name: GetUserByAPIKey :one
SELECT id, created_at, updated_at, name, apikey, username, password FROM users WHERE apikey = $1
`

func (q *Queries) GetUserByAPIKey(ctx context.Context, apikey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAPIKey, apikey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Apikey,
		&i.Username,
		&i.Password,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, created_at, updated_at, name, apikey, username, password FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Apikey,
		&i.Username,
		&i.Password,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, created_at, updated_at, name, apikey, username, password FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Apikey,
		&i.Username,
		&i.Password,
	)
	return i, err
}
