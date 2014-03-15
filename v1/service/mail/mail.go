// Adapted from Google Appengine mail package.
package mail

import (
	"bytes"
	"log"
	"net/smtp"
	"text/template"

	"github.com/hackedu/backend/v1/helper"
)

var auth smtp.Auth

func init() {
	auth = smtp.PlainAuth(
		"",
		helper.GetConfig("SMTP_USERNAME"),
		helper.GetConfig("SMTP_PASSWORD"),
		helper.GetConfig("SMTP_HOST"),
	)
}

type Message struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

const emailTemplate = `From: {{.Sender}}
To: {{range .To}}{{.}},{{end}}
Subject: {{.Subject}}

{{.Body}}
`

// Send sends an email message.
func Send(msg *Message) error {
	var doc bytes.Buffer

	t := template.New("emailTemplate")
	t, err := t.Parse(emailTemplate)
	if err != nil {
		log.Println("Error parsing mail template", err)
		return err
	}
	err = t.Execute(&doc, msg)
	if err != nil {
		log.Println("Error executing mail template", err)
		return err
	}

	err = smtp.SendMail(
		helper.GetConfig("SMTP_HOST")+":"+helper.GetConfig("SMTP_PORT"),
		auth,
		msg.Sender,
		msg.To,
		doc.Bytes(),
	)

	if err != nil {
		return err
	}

	return nil
}
