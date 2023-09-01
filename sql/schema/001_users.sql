-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULl,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);


-- +goose Down

DROP TABLE users;