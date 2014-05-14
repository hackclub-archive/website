package model

import (
	"encoding/json"
	"errors"
	"regexp"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
)

var (
	// ErrInvalidFirstName is returned when the user's first name is invalid.
	ErrInvalidFirstName = errors.New("invalid first name")
	// ErrInvalidLastName is returned when the user's last name is invalid.
	ErrInvalidLastName = errors.New("invalid last name")
	// ErrInvalidEmail is returned when the user's email is invalid.
	ErrInvalidEmail = errors.New("invalid email address")
	// ErrInvalidPassword is returned when the user's password is invalid.
	ErrInvalidPassword = errors.New("invalid password")
)

var regexpEmail = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

// User represents a user of hackEDU.
type User struct {
	ID        int64     `db:"id"         json:"id"`
	Created   time.Time `db:"created"    json:"created"`
	Updated   time.Time `db:"updated"    json:"updated"`
	FirstName string    `db:"first_name" json:"firstName"`
	LastName  string    `db:"last_name"  json:"lastName"`
	Email     string    `db:"email"      json:"email"`
	GitHub    string    `db:"github"     json:"github"`
	Twitter   string    `db:"twitter"    json:"twitter"`
	Password  string    `db:"password"   json:"-"`
}

type intermediateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	GitHub    string `json:"github"`
	Twitter   string `json:"twitter"`
	Password  string `json:"password"`
}

// NewUser creates a new user from provided JSON. It unmarshales the JSON,
// validates the fields, generates a hash from the provided password string
// using bcrypt, and then returns the created user.
//
// NewUser does not save the user to the database.
func NewUser(jsonData []byte) (*User, error) {
	var iU intermediateUser
	if err := json.Unmarshal(jsonData, &iU); err != nil {
		return nil, err
	}

	if err := iU.validate(); err != nil {
		return nil, err
	}

	b, err := bcrypt.GenerateFromPassword([]byte(iU.Password),
		bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		FirstName: iU.FirstName,
		LastName:  iU.LastName,
		Email:     iU.Email,
		GitHub:    iU.GitHub,
		Twitter:   iU.Twitter,
		Password:  string(b),
	}

	return &user, nil
}

func (u *intermediateUser) validate() error {
	switch {
	case len(u.FirstName) == 0:
		return ErrInvalidFirstName
	case len(u.FirstName) >= 255:
		return ErrInvalidFirstName
	case len(u.LastName) == 0:
		return ErrInvalidLastName
	case len(u.LastName) >= 255:
		return ErrInvalidLastName
	case len(u.Email) >= 255:
		return ErrInvalidEmail
	case regexpEmail.MatchString(u.Email) == false:
		return ErrInvalidEmail
	case len(u.Password) < 6:
		return ErrInvalidPassword
	case len(u.Password) > 256:
		return ErrInvalidPassword
	default:
		return nil
	}
}
