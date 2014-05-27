
-- +goose Up
ALTER TABLE users ALTER COLUMN created TYPE timestamp;
ALTER TABLE users ALTER COLUMN updated TYPE timestamp;
ALTER TABLE schools ALTER COLUMN created TYPE timestamp;
ALTER TABLE schools ALTER COLUMN updated TYPE timestamp;


-- +goose Down
ALTER TABLE users ALTER COLUMN created TYPE date;
ALTER TABLE users ALTER COLUMN updated TYPE date;
ALTER TABLE schools ALTER COLUMN created TYPE date;
ALTER TABLE schools ALTER COLUMN updated TYPE date;
