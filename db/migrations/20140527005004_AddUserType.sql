
-- +goose Up
ALTER TABLE users ADD COLUMN type integer not null default 2;


-- +goose Down
ALTER TABLE users DROP COLUMN type;

