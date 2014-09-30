package database

import (
	"time"

	"github.com/hackedu/hackedu/v1/club"
	"github.com/hackedu/hackedu/v1/user"
)

const clubGetByIDStmt = `SELECT id, created, updated, school_id, name FROM
clubs WHERE id=$1`

const clubGetByIDForUser = `
SELECT c.id, c.created, c.updated, c.school_id, c.name
FROM   clubs       c
JOIN   users_clubs uc USING (id)
WHERE  c.id = $1 AND uc.user_id = $2`

const clubGetAllStmt = `SELECT id, created, updated, school_id, name FROM
clubs ORDER BY id`

const clubGetAllForUser = `
SELECT c.id, c.created, c.updated, c.school_id, c.name
FROM   clubs       c
JOIN   users_clubs uc USING (id)
WHERE  uc.user_id = $1`

const clubCreateStmt = `INSERT INTO clubs (created, updated, school_id, name)
VALUES ($1, $2, $3, $4) RETURNING id`

const clubCreateRelationshipStmt = `INSERT INTO users_clubs (user_id,
club_id) VALUES ($1, $2)`

// GetClub gets a club from the database with the provided ID
func GetClub(id int64) (*club.Club, error) {
	c := club.Club{}
	row := db.QueryRow(clubGetByIDStmt, id)
	if err := row.Scan(&c.ID, &c.Created, &c.Updated, &c.SchoolID,
		&c.Name); err != nil {
		return nil, err
	}
	return &c, nil
}

// GetClubForUser gets a club from the database with the provided ID that the
// given user ID has an association with.
func GetClubForUser(clubID, userID int64) (*club.Club, error) {
	c := club.Club{}
	row := db.QueryRow(clubGetByIDForUser, clubID, userID)
	if err := row.Scan(&c.ID, &c.Created, &c.Updated, &c.SchoolID,
		&c.Name); err != nil {
		return nil, err
	}
	return &c, nil
}

// GetClubs gets all of the clubs from the database ordered by id.
func GetClubs() ([]*club.Club, error) {
	clubs := []*club.Club{}
	rows, err := db.Query(clubGetAllStmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := club.Club{}
		if err := rows.Scan(&c.ID, &c.Created, &c.Updated, &c.SchoolID,
			&c.Name); err != nil {
			return nil, err
		}

		clubs = append(clubs, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clubs, nil
}

// GetClubsForUsers returns all of the clubs that the provided user has a
// relationship with.
func GetClubsForUser(userID int64) ([]*club.Club, error) {
	clubs := []*club.Club{}
	rows, err := db.Query(clubGetAllForUser, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := club.Club{}
		if err := rows.Scan(&c.ID, &c.Created, &c.Updated, &c.SchoolID,
			&c.Name); err != nil {
			return nil, err
		}

		clubs = append(clubs, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clubs, nil
}

// SaveClub saves the provided club to the database. If the club is a new
// club, then the club.Created field is set to the current time. The
// club.Updated field is set to the current time regardless.
func SaveClub(c *club.Club, u *user.User) error {
	if c.ID == 0 {
		c.Created = time.Now()
	}
	c.Updated = time.Now()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(clubCreateStmt, c.Created, c.Updated, c.SchoolID, c.Name)
	if err := row.Scan(&c.ID); err != nil {
		return err
	}

	_, err = tx.Exec(clubCreateRelationshipStmt, u.ID, c.ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// AddUserToClub creates a new user association with a club.
func AddUserToClub(userID, clubID int64) error {
	_, err := db.Exec(clubCreateRelationshipStmt, userID, clubID)
	if err != nil {
		return err
	}
	return nil
}
