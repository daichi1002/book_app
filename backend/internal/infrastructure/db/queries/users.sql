-- name: GetUser :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CreateUser :exec
INSERT INTO users (name, email) VALUES (?, ?);

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
