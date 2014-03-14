
-- +goose Up
CREATE TABLE Users (
  Id serial not null primary key,
  CreatedAt date not null,
  FirstName text not null,
  LastName text not null,
  Email text not null,
  GitHub text,
  Twitter text,
  HashedPassword bytea not null
);

CREATE TABLE Application (
  Id serial not null primary key,
  UserId integer not null,
  CreatedAt date not null,
  HighSchool text not null,
  InterestingProject text not null,
  SystemHacked text not null,
  Passion text not null,
  Story text not null,
  Why text not null
);

ALTER TABLE Application
  ADD FOREIGN KEY (UserId) REFERENCES Users(Id)


-- +goose Down
DROP TABLE Users;
DROP TABLE Application;
