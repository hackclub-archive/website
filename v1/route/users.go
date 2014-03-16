package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"code.google.com/p/go.crypto/scrypt"

	"github.com/coopernurse/gorp"
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
		log.Println(err)
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
		To:      []string{"Zach Latta <zach@zachlatta.com>"},
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
		log.Println(err)
		return http.StatusInternalServerError, "Could not send email"
	}

	msg = &mail.Message{
		Sender: "Zach Latta <zach@zachlatta.com>",
		To: []string{
			fmt.Sprintf("%s %s <%s>", user.FirstName, user.LastName, user.Email),
		},
		Subject: "hackEDU Application",
		Body: `Hey ` + user.FirstName + `!

Thanks for applying for hackEDU. We've received your application and you can
expect to hear from us shortly. If you have any questions, please don't
hesitate to email me at zach@zachlatta.com.

Best regards,
Zach Latta
`,
	}

	if err := mail.Send(msg); err != nil {
		return http.StatusInternalServerError, "Could not send email"
	}

	json, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, "Error while creating user."
	}

	return http.StatusOK, string(json)
}

func hashPassword(password, salt string) ([]byte, error) {
	return scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
}
