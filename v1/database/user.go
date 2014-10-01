package database

import (
	"database/sql"
	"time"

	"github.com/hackedu/hackedu/v1/user"
)

const userGetByIDStmt = `SELECT id, created, updated, first_name, last_name,
email, type, github, twitter, password FROM users WHERE id = $1`

const userGetByEmailStmt = `SELECT id, created, updated, first_name, last_name,
email, type, github, twitter, password FROM users WHERE email ilike $1`

const userCreateStmt = `INSERT INTO users (created, updated, first_name,
last_name, email, type, github, twitter, password) VALUES ($1, $2, $3, $4, $5,
$6, $7, $8, $9) RETURNING id`

// GetUser gets the user from the database with the provided ID.
func GetUser(id int64) (*user.User, error) {
	u := new(user.User)
	row := db.QueryRow(userGetByIDStmt, id)
	if err := row.Scan(&u.ID, &u.Created, &u.Updated, &u.FirstName, &u.LastName,
		&u.Email, &u.Type, &u.GitHub, &u.Twitter, &u.Password); err != nil {
		return nil, err
	}
	return u, nil
}

// GetUserByEmail gets the user from the database with the provided email.
func GetUserByEmail(email string) (*user.User, error) {
	u := new(user.User)
	row := db.QueryRow(userGetByEmailStmt, email)
	if err := row.Scan(&u.ID, &u.Created, &u.Updated, &u.FirstName, &u.LastName,
		&u.Email, &u.Type, &u.GitHub, &u.Twitter, &u.Password); err != nil {
		return nil, err
	}
	return u, nil
}

// SaveUser saves the provided user to the database. If the user is a new
// user, then the user.Created field is set to the current time. The
// user.Updated field is set to the current time regardless.
//
// If the user is a new user, then SaveUser also verifies that the user's
// email is unique and returns user.ErrInvalidEmail accordingly.
func SaveUser(u *user.User) error {
	if u.ID == 0 {
		// TODO: Should do this in the user.validate() method, but with the
		// current architecture of the application, that causes an import cycle.
		//
		// Check if email is unique.
		_, err := GetUserByEmail(u.Email)
		if err == nil {
			return user.ErrInvalidUserEmail
		} else if err != sql.ErrNoRows && err != nil {
			return err
		}

		u.Created = time.Now()
	}
	u.Updated = time.Now()

	rows, err := db.Query(userCreateStmt, u.Created, u.Updated, u.FirstName,
		u.LastName, u.Email, u.Type, u.GitHub, u.Twitter, u.Password)
	if err != nil {
		return err
	}
	for rows.Next() {
		if err := rows.Scan(&u.ID); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
