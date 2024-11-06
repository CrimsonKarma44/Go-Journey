-- name: CreateBlog :execresult
INSERT INTO blog (title, content, category, tags, createdAt, updatedAt)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateBlog :exec
UPDATE blog
SET title = ?,
    content = ?,
    category = ?,
    tags = ?,
    updatedAt = ?
WHERE id = ?;

-- name: ViewAllBlog :many
select *
from blog;

-- name: ViewSingleBlog :one
select *
from blog
WHERE id = ?;

-- name: DeleteBlog :exec
DELETE
FROM blog
WHERE id = ?;

-- name: TermBlogSearch :many
SELECT *
FROM blog
WHERE title LIKE CONCAT('%', ?, '%')
   OR content LIKE CONCAT('%', ?, '%')
   OR category LIKE CONCAT('%', ?, '%');
