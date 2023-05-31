-- name: CreateUser :one
INSERT INTO users (role_id, email, username, password, first_name, last_name, is_protected) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUserName :one
SELECT * FROM users WHERE username = $1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: UpdateUser :one
UPDATE users SET role_id = $2, email = $3, username = $4, password = $5, first_name = $6, last_name = $7, is_protected = $8, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;