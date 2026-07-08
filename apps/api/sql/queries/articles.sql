-- name: CreateArticle :one
INSERT INTO articles (author_id, title, content, status)
VALUES ($1, $2, $3, $4)
RETURNING id, author_id, title, content, status, synced_to_graph, created_at, updated_at;

-- name: GetArticleByID :one
SELECT id, author_id, title, content, status, synced_to_graph, created_at, updated_at
FROM articles
WHERE id = $1;

-- name: UpdateArticle :one
UPDATE articles
SET title = $2, content = $3, status = $4
WHERE id = $1
RETURNING id, author_id, title, content, status, synced_to_graph, created_at, updated_at;
