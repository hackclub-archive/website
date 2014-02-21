package hackedu

import (
	"encoding/json"
	"errors"
	"net/http"

	"appengine"
	"appengine/datastore"
	"appengine/mail"
)

type ApplicationForm struct {
	Application Application `json:"application,omitempty"`
	User        User        `json:"user,omitempty"`
}

func Apply(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	decoder := json.NewDecoder(r.Body)
	var applicationForm ApplicationForm
	err := decoder.Decode(&applicationForm)
	if err != nil {
		serveError(c, w, err)
		return
	}

	user := applicationForm.User
	application := applicationForm.Application

	userKey, err := RegisterUser(c, &user)
	if err != nil {
		serveError(c, w, err)
		return
	}

	applicationKey, err := ConstructApplication(c, &application)
	if err != nil {
		serveError(c, w, err)
		return
	}

	user.Application = applicationKey

	_, err = datastore.Put(c, userKey, &user)
	if err != nil {
		serveError(c, w, err)
		return
	}

	msg := &mail.Message{
		Sender:  user.FirstName + " " + user.LastName + " <" + user.Email + ">",
		To:      []string{"zach@zachlatta.com"},
		Subject: "hackEDU Application",
		Body: `
# User Information

Name: ` + user.FirstName + ` ` + user.LastName + `
Email: ` + user.Email + `
GitHub: ` + user.GitHub + `
Twitter: ` + user.Twitter + `

# Application Information

High School: ` + application.HighSchool + `

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

	if err := mail.Send(c, msg); err != nil {
		serveError(c, w, errors.New("Could not send email"))
		return
	}

	msg = &mail.Message{
		Sender:  "Zach Latta <zach@zachlatta.com>",
		To:      []string{user.Email},
		Subject: "hackEDU Application",
		Body: `
Hey ` + user.FirstName + `!

Thanks for applying for hackEDU. We've received your application and you can
expect to hear from us shortly. If you have any questions, please don't
hesitate to email me at zach@zachlatta.com.

Best regards,
Zach Latta
`,
	}

	if err := mail.Send(c, msg); err != nil {
		serveError(c, w, errors.New("Could not send email"))
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		serveError(c, w, err)
		return
	}
}
