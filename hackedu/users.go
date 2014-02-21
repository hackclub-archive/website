package hackedu

import (
	"errors"
	"regexp"
	"time"

	"code.google.com/p/go.crypto/scrypt"

	"appengine"
	"appengine/datastore"
)

const emailRegex = ".+\\@.+\\..+"

type User struct {
	CreatedAt      time.Time      `json:"created_at,omitempty"`
	FirstName      string         `json:"first_name,omitempty"`
	LastName       string         `json:"last_name,omitempty"`
	Email          string         `json:"email,omitempty"`
	Password       string         `json:"password,omitempty"`
	PasswordVerify string         `json:"password_verify,omitempty"`
	HashedPassword []byte         `json:"-"`
	Application    *datastore.Key `json:"application,omitempty"`
}

func RegisterUser(c appengine.Context, user *User) (*datastore.Key, error) {

	if len(user.Password) < 6 {
		return nil, errors.New("Your password must be at least 6 characters long")
	}

	if user.Password != user.PasswordVerify {
		return nil, errors.New("Password does not match password verify.")
	}

	if !(len(user.FirstName) > 0) {
		return nil, errors.New("A first name is required.")
	}

	if !(len(user.LastName) > 0) {
		return nil, errors.New("A last name is required.")
	}

	if match, _ := regexp.MatchString(emailRegex, user.Email); !match {
		return nil, errors.New("A valid email is required.")
	}

	// TODO: Check if email has already been taken

	user.CreatedAt = time.Now()

	hashedPassword, err := scrypt.Key([]byte(user.Password),
		[]byte(user.CreatedAt.String()+user.Email+"douglasadams42"),
		16384, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	user.HashedPassword = hashedPassword
	user.Password = ""
	user.PasswordVerify = ""

	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "user", nil),
		user)
	if err != nil {
		return key, err
	}

	return key, nil
}
