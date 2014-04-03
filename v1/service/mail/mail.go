// Adapted from Google Appengine mail package.
package mail

import (
	"bytes"
	"errors"
	"io/ioutil"
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

// An email message. Include either the Body or Template field. If both are
// included, only Template will be used.
type Message struct {
	Sender   string
	To       []string
	Subject  string
	Body     string
	Template string      // Path of template to load
	Context  interface{} // required if Template is used
}

const emailTemplate = `From: {{.Sender}}
To: {{range .To}}{{.}},{{end}}
Subject: {{.Subject}}

{{.Body}}
`

const tmplPath = "v1/template/email/"

// Send sends an email message.
func Send(msg *Message) error {
	if msg.Template != "" {
		tmplBytes, err := ioutil.ReadFile(tmplPath + msg.Template)
		if err != nil {
			return err
		}

		t := template.Must(template.New("emailBody").Parse(string(tmplBytes)))

		var doc bytes.Buffer
		err = t.Execute(&doc, msg.Context)
		if err != nil {
			return err
		}

		msg.Body = string(doc.Bytes())
	}

	return send(msg)
}

func send(msg *Message) error {
	var doc bytes.Buffer

	t := template.New("emailTemplate")
	t, err := t.Parse(emailTemplate)
	if err != nil {
		return errors.New("Error parsing mail template: " + err.Error())
	}
	err = t.Execute(&doc, msg)
	if err != nil {
		return errors.New("Error executing mail template: " + err.Error())
	}

	err = smtp.SendMail(
		helper.GetConfig("SMTP_HOST")+":"+helper.GetConfig("SMTP_PORT"),
		auth,
		msg.Sender,
		msg.To,
		doc.Bytes(),
	)
	if err != nil {
		return errors.New("Error sending email: " + err.Error())
	}

	return nil
}
