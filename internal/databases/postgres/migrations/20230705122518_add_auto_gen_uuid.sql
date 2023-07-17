-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE users ALTER COLUMN id SET DEFAULT gen_random_uuid(); 

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
ALTER TABLE users ALTER COLUMN id DROP DEFAULT;
