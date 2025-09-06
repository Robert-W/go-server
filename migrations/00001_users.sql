-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY,
  name TEXT NOT NULL,
  created TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,
  last_updated TIMESTAMP default CURRENT_TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
