
-- +goose Up
CREATE TABLE User (
  Id integer not null primary key autoincrement,
  CreatedAt datetime not null,
  FirstName text not null,
  LastName text not null,
  Email text not null,
  GitHub text,
  Twitter text,
  HashedPassword varbinary(1024) not null
);

CREATE TABLE Application (
  Id integer not null primary key autoincrement,
  UserId integer not null,
  CreatedAt datetime not null,
  HighSchool text not null,
  InterestingProject text not null,
  SystemHacked text not null,
  Passion text not null,
  Story text not null,
  Why text not null
);

ALTER TABLE Application
  ADD FOREIGN KEY (UserId) REFERENCES User(Id)


-- +goose Down
DROP TABLE User;
DROP TABLE Application;
