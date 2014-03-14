
-- +goose Up
CREATE TABLE School (
  Id integer not null primary key autoincrement,
  CreatedAt datetime,
  Name text,
  Latitude real,
  Longitude real 
);


-- +goose Down
DROP TABLE School;
