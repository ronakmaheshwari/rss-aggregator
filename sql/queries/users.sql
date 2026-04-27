-- name: CreateUser :one
INSERT INTO users (id, name, email, created_at, updated_at)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetUsers :many
SELECT id, name, email, created_at, updated_at FROM users;

-- name: GetUserByEmail :one
SELECT id, name, email, created_at, updated_at FROM users WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET name = $1,
    updated_at = $2
WHERE email = $3
RETURNING id, name, email, created_at, updated_at;