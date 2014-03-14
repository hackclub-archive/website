
-- +goose Up
CREATE TABLE School (
  Id serial not null primary key,
  CreatedAt date not null,
  Name text not null,
  Latitude real not null,
  Longitude real not null 
);


-- +goose Down
DROP TABLE School;
