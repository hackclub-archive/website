
-- +goose Up
ALTER TABLE School
  ADD COLUMN Website text;

-- +goose Down
ALTER TABLE School
  DROP COLUMN Website;
