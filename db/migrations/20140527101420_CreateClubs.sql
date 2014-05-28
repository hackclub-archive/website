
-- +goose Up
CREATE TABLE clubs (
  id serial not null primary key,
  created date not null,
  updated date not null,
  school_id integer references schools(id) ON UPDATE CASCADE ON DELETE CASCADE,
  name text not null
);

CREATE TABLE users_clubs (
  id serial not null primary key,
  user_id integer references users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  club_id integer references clubs(id) ON UPDATE CASCADE ON DELETE CASCADE
);


-- +goose Down
DROP TABLE users_clubs;
DROP TABLE clubs;

