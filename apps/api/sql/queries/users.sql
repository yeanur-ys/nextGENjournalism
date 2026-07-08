-- name: CreateUser :one
INSERT INTO users (email, password_hash, display_name, role)
VALUES ($1, $2, $3, $4)
RETURNING id, email, display_name, role, bio, verification, created_at;

-- name: GetUserByEmail :one
SELECT id, email, display_name, role, bio, verification, created_at, password_hash
FROM users
WHERE email = $1;
