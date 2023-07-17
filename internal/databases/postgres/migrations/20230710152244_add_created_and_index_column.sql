-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE users ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT(now() AT TIME ZONE 'utc');
CREATE INDEX index_user_pagination ON users (created_at, id);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP INDEX index_user_pagination;
ALTER TABLE users DROP COLUMN created_at;
