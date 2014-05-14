package model

import (
	"encoding/json"
	"errors"
	"regexp"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
)

var (
	ErrInvalidFirstName = errors.New("Invalid first name")
	ErrInvalidLastName  = errors.New("Invalid last name")
	ErrInvalidEmail     = errors.New("Invalid email address")
	ErrInvalidPassword  = errors.New("Invalid password")
)

var RegexpEmail = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

type User struct {
	ID        int64     `json:"id"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	GitHub    string    `json:"github"`
	Twitter   string    `json:"string"`
	Password  string    `json:"-"`
}

type intermediateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	GitHub    string `json:"github"`
	Twitter   string `json:"twitter"`
	Password  string `json:"password"`
}

func NewUser(jsonData []byte) (*User, error) {
	var iU intermediateUser
	if err := json.Unmarshal(jsonData, &iU); err != nil {
		return nil, err
	}

	if err := iU.validate(); err != nil {
		return nil, err
	}

	b, err := bcrypt.GenerateFromPassword([]byte(iU.Password), bcrypt.DefaultCost)
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
	case RegexpEmail.MatchString(u.Email) == false:
		return ErrInvalidEmail
	case len(u.Password) < 6:
		return ErrInvalidPassword
	case len(u.Password) > 256:
		return ErrInvalidPassword
	default:
		return nil
	}
}
