
-- +goose Up
CREATE TYPE user_type AS ENUM ('admin', 'organizer', 'student');

ALTER TABLE users DROP COLUMN type;
ALTER TABLE users
  ADD COLUMN type user_type NOT NULL DEFAULT 'student';


-- +goose Down
ALTER TABLE users DROP COLUMN type;
ALTER TABLE users ADD COLUMN type integer not null default 2;
DROP TYPE user_type;

