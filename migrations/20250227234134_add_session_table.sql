-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sessions(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  user_agent TEXT,
  ip_address TEXT,
  created_at TIMESTAMP NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd


