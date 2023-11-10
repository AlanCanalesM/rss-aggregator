-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds SET last_fetched_at = NOW(), updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: FeedsNotFollowedByUser :many
SELECT f.*
FROM public.feeds f
LEFT JOIN public.feed_follows ff ON f.id = ff.feed_id AND ff.user_id = $1
WHERE ff.id IS NULL;
