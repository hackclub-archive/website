
-- +goose Up
CREATE TABLE users (
  id serial not null primary key,
  created date not null,
  updated date not null,
  first_name text not null,
  last_name text not null,
  email text not null unique,
  github text,
  twitter text,
  password text not null
);


-- +goose Down
DROP TABLE users;
