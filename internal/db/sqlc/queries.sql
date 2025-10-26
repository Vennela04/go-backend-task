-- Users table queries

-- name: CreateUser :one
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING id, name, dob;

-- name: GetUserByID :one
SELECT id, name, dob
FROM users
WHERE id = $1;

-- name: GetAllUsers :many
SELECT id, name, dob
FROM users;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $2, dob = $3
WHERE id = $1
RETURNING id, name, dob;
