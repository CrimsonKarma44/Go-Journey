-- name: GetAuthorData :one
SELECT id, username, password
FROM users
WHERE username = ? LIMit 1;

-- name: GetAuthor :one
SELECT *
FROM users
WHERE id = ? LIMIT 1;

-- name: ListAuthors :many
SELECT *
FROM users
ORDER BY name;

-- name: CreateAuthor :execresult
INSERT INTO users (username, email, password)
VALUES (?, ?, ?);


-- name: DeleteAuthor :exec
DELETE
FROM users
WHERE id = ?;