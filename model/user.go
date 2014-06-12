package model

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"time"

	"crypto/rand"

	"code.google.com/p/go.crypto/bcrypt"
)

const (
	// UserAdmin is the user's type when the user is an admin
	UserAdmin = "admin"
	// UserAdmin is the user's type when the user is a club organizer
	UserOrganizer = "organizer"
	// UserAdmin is the user's type when the user is a student in a club
	UserStudent = "student"
)

var (
	// ErrInvalidUserFirstName is returned when the user's first name is invalid.
	ErrInvalidUserFirstName = errors.New("invalid first name")
	// ErrInvalidUserLastName is returned when the user's last name is invalid.
	ErrInvalidUserLastName = errors.New("invalid last name")
	// ErrInvalidUserEmail is returned when the user's email is invalid.
	ErrInvalidUserEmail = errors.New("invalid email address")
	// ErrInvalidUserType is returned when the user's type is invalid.
	ErrInvalidUserType = errors.New("invalid type")
	// ErrInvalidUserPassword is returned when the user's password is invalid.
	ErrInvalidUserPassword = errors.New("invalid password")
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
	Type      string    `db:"type"       json:"type"`
	GitHub    string    `db:"github"     json:"github"`
	Twitter   string    `db:"twitter"    json:"twitter"`
	Password  string    `db:"password"   json:"-"`
}

// RequestUser represents a user of hackEDU as passed by the frontend.
// RequestUser will need to be transformed into a User to be stored into the
// database.
type RequestUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Type      string `json:"type"`
	GitHub    string `json:"github"`
	Twitter   string `json:"twitter"`
	Password  string `json:"password"`
}

// NewUser creates a new user from provided JSON reader. It decodes the JSON,
// validates the fields, generates a hash from the provided password string
// using bcrypt, and then returns the created user.
//
// NewUser does not save the user to the database.
func NewUser(jsonReader io.Reader) (*User, error) {
	var rU RequestUser
	if err := json.NewDecoder(jsonReader).Decode(&rU); err != nil {
		return nil, err
	}

	if err := rU.validate(); err != nil {
		return nil, err
	}

	b, err := bcrypt.GenerateFromPassword([]byte(rU.Password),
		bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		FirstName: rU.FirstName,
		LastName:  rU.LastName,
		Email:     rU.Email,
		Type:      rU.Type,
		GitHub:    rU.GitHub,
		Twitter:   rU.Twitter,
		Password:  string(b),
	}

	return &user, nil
}

// NewUserGeneratePassword creates a new user from the provided JSON reader.
// It decodes the JSON, validates the fields, generates a password and hashes
// it using bcrypt, and then returns the created user.
//
// NewUserGeneratePassword does not save the user to the database.
func NewUserGeneratePassword(jsonReader io.Reader) (*User, error) {
	var rU RequestUser
	if err := json.NewDecoder(jsonReader).Decode(&rU); err != nil {
		return nil, err
	}

	b, err := bcrypt.GenerateFromPassword([]byte(generatePassword(16)),
		bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if err := rU.validate(); err != nil {
		return nil, err
	}

	user := User{
		FirstName: rU.FirstName,
		LastName:  rU.LastName,
		Email:     rU.Email,
		Type:      rU.Type,
		GitHub:    rU.GitHub,
		Twitter:   rU.Twitter,
		Password:  string(b),
	}

	return &user, nil
}

func (u *RequestUser) validate() error {
	switch {
	case len(u.FirstName) == 0:
		return ErrInvalidUserFirstName
	case len(u.FirstName) >= 255:
		return ErrInvalidUserFirstName
	case len(u.LastName) == 0:
		return ErrInvalidUserLastName
	case len(u.LastName) >= 255:
		return ErrInvalidUserLastName
	case len(u.Email) >= 255:
		return ErrInvalidUserEmail
	case regexpEmail.MatchString(u.Email) == false:
		return ErrInvalidUserEmail
	case !(u.Type == UserAdmin || u.Type == UserOrganizer ||
		u.Type == UserStudent):
		return ErrInvalidUserType
	case len(u.Password) < 6:
		return ErrInvalidUserPassword
	case len(u.Password) > 256:
		return ErrInvalidUserPassword
	default:
		return nil
	}
}

// ComparePassword compares the supplied password to the user password stored
// in the database.
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func generatePassword(length int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
