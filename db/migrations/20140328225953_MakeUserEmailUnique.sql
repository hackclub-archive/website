
-- +goose Up
ALTER TABLE Users
  ADD CONSTRAINT users_email_key UNIQUE (Email);


-- +goose Down
ALTER TABLE Users
  DROP CONSTRAINT users_email_key;
