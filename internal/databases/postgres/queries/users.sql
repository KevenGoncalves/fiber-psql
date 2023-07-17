-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (name,email,last_name) VALUES ($1, $2, $3) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUser :one 
UPDATE users SET name = $2, email = $3, last_name = $4 WHERE id = $1 RETURNING *;

