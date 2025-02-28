-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
firstname TEXT NOT NULL,
lastname TEXT NOT NULL,
email TEXT NOT NULL,
deleted_at TIMESTAMP NULL,
verified BOOL NOT NULL,
password TEXT NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd


