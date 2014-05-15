
-- +goose Up
CREATE TABLE schools (
  id serial not null primary key,
  created date not null,
  updated date not null,
  name text not null,
  website text,
  latitude real not null,
  longitude real not null
);


-- +goose Down
DROP TABLE schools;

