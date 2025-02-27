-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN refresh_token TEXT NULL;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN refresh_token;
-- +goose StatementEnd


