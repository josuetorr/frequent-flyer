-- +goose Up
-- +goose StatementBegin
-- CREATE EXTENSIONS IF NOT EXISTS "pgcrypto";
CREATE TABLE IF NOT EXISTS users(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
firstname TEXT NOT NULL,
lastname TEXT NOT NULL,
email TEXT NOT NULL,
verified BOOL NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd


