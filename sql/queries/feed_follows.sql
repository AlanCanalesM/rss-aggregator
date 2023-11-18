-- name: CreatedFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id, feed_name)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE feed_id = $1 AND user_id = $2;