-- name: CreateBlog :execresult
INSERT INTO blog (title, content, category, tags, createdAt, updatedAt)
VALUES (?, ?, ?, ?, ?, ?);

-- name: DeleteBlog :exec
DELETE
FROM blog
WHERE id = ?;
