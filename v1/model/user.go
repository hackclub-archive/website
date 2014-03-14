package model

import (
	"net/http"
	"regexp"
	"time"

	"github.com/martini-contrib/binding"
)

const emailRegex = ".+\\@.+\\..+"

type User struct {
	Id             int          `json:"userId"`
	CreatedAt      time.Time    `json:"createdAt"`
	FirstName      string       `json:"firstName" binding:"required"`
	LastName       string       `json:"lastName" binding:"required"`
	Email          string       `json:"email" binding:"required"`
	GitHub         string       `json:"github,omitempty"`
	Twitter        string       `json:"twitter,omitempty"`
	Password       string       `json:"password,omitempty" binding:"required" db:"-"`
	PasswordVerify string       `json:"passwordVerify,omitempty" binding:"required" db:"-"`
	HashedPassword []byte       `json:"-"`
	Application    *Application `json:"application,omitempty" binding:"required" db:"-"`
}

func (u User) Validate(errors *binding.Errors, r *http.Request) {
	if len(u.Password) < 6 {
		errors.Fields["Password"] =
			"Your password must be at least 6 characters long"
	}

	if u.Password != u.PasswordVerify {
		errors.Fields["PasswordVerify"] =
			"Password does not match password verify."
	}

	if match, _ := regexp.MatchString(emailRegex, u.Email); !match {
		errors.Fields["Email"] = "A valid email is required."
	}
}
