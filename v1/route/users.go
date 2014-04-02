package route

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"code.google.com/p/go.crypto/scrypt"

	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/hackedu/backend/v1/model"
	"github.com/hackedu/backend/v1/service/mail"
)

func AddUser(user model.User, db gorp.SqlExecutor, log *log.Logger) (int, string) {
	application := user.Application
	user.Application = nil

	user.CreatedAt = time.Now()
	application.CreatedAt = time.Now()

	salt := user.CreatedAt.String() + user.LastName + user.Email

	hashedPassword, err := hashPassword(user.Password, salt)
	if err != nil {
		log.Println("Error hashing password", err)
		return http.StatusInternalServerError, "Error while creating user."
	}

	user.Password = ""
	user.PasswordVerify = ""
	user.HashedPassword = hashedPassword

	// TODO: Figure out why this isn't doing anything.
	db.Insert(&user)

	application.UserId = user.Id

	db.Insert(application)

	msg := &mail.Message{
		Sender:  user.FirstName + " " + user.LastName + " <" + user.Email + ">",
		To:      []string{"Zach Latta <zach@hackedu.us>"},
		Subject: "hackEDU Application",
		Body: `# User Information

Name: ` + user.FirstName + ` ` + user.LastName + `
Email: ` + user.Email + `
GitHub: ` + user.GitHub + `
Twitter: ` + user.Twitter + `

# Application

## High School

` + application.HighSchool + `

## Interesting Project

` + application.InterestingProject + `

## System Hacked

` + application.SystemHacked + `

## Passion

` + application.Passion + `

## Story 

` + application.Story + `

## Why

` + application.Why + `

`,
	}

	if err := mail.Send(msg); err != nil {
		log.Println("Could not send email", err)
		return http.StatusInternalServerError, "Could not send email"
	}

	msg = &mail.Message{
		Sender: "Zach Latta <zach@hackedu.us>",
		To: []string{
			fmt.Sprintf("%s %s <%s>", user.FirstName, user.LastName, user.Email),
		},
		Subject: "hackEDU Application",
		Body: `Hey ` + user.FirstName + `!

Thanks for applying for hackEDU. We've received your application and you can
expect to hear from us shortly. If you have any questions, please don't
hesitate to email me at zach@hackedu.us.

Best regards,
Zach Latta
`,
	}

	if err := mail.Send(msg); err != nil {
		log.Println("Could not send email", err)
		return http.StatusInternalServerError, "Could not send email"
	}

	json, err := json.Marshal(user)
	if err != nil {
		log.Println("Could not marshal user to JSON", err)
		return http.StatusInternalServerError, "Error while creating user."
	}

	return http.StatusCreated, string(json)
}

func hashPassword(password, salt string) ([]byte, error) {
	return scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
}

// GetUser retrieves a user from the database. If the id passed is an int,
// the user is retrieved by id. Else, the user is retrieved by email.
func GetUser(db gorp.SqlExecutor, params martini.Params,
	log *log.Logger) (int, string) {
	var user *model.User
	rawID := params["id"]

	id, err := strconv.Atoi(rawID)
	if err != nil {
		var statusCode int
		user, err, statusCode = getUserByEmail(rawID, db)
		if err != nil {
			log.Printf("Error retrieving user with email %s: %v", rawID, err)
			return statusCode, "Error retrieving user"
		}
	} else {
		var statusCode int
		user, err, statusCode = getUserByID(id, db)
		if err != nil {
			log.Printf("Error retrieving user with id %d: %v", id, err)
			return statusCode, "Error retrieving user"
		}
	}

	json, err := json.Marshal(user)
	if err != nil {
		log.Println("Could not marshal user to JSON", err)
		return http.StatusInternalServerError, "Error while creating user."
	}

	return http.StatusOK, string(json)
}

func getUserByID(id int, db gorp.SqlExecutor) (*model.User, error, int) {
	obj, err := db.Get(model.User{}, id)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	if obj == nil {
		return nil, errors.New("user not found"), http.StatusNotFound
	}

	return obj.(*model.User), nil, http.StatusOK
}

func getUserByEmail(email string,
	db gorp.SqlExecutor) (*model.User, error, int) {
	user := model.User{}

	err := db.SelectOne(&user,
		fmt.Sprintf("SELECT * FROM Users WHERE Email ilike '%s'", email))
	if err != nil {
		return nil, err, http.StatusNotFound
	}

	return &user, nil, http.StatusOK
}
