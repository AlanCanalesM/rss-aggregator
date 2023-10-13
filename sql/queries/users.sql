-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, apikey, username, password) 
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'), $5, $6) RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE apikey = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;