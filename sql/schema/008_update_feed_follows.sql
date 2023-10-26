-- +goose Up
ALTER TABLE feed_follows
ADD COLUMN feed_name TEXT NOT NULL DEFAULT '';
UPDATE feed_follows AS ff
SET feed_name = f.name
FROM feeds AS f
WHERE ff.feed_id = f.id;

-- +goose Down
ALTER TABLE feed_follows
DROP COLUMN feed_name;
```