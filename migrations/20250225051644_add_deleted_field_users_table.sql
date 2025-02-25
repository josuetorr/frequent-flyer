-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN deleted_at TIMESTAMP NULL;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN deleted_at;

;
-- +goose StatementEnd


