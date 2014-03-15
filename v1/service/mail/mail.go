// Adapted from Google Appengine mail package.
package mail

import (
	"net/smtp"
	"os"
)

var auth smtp.Auth

func init() {
	auth = smtp.PlainAuth(
		"",
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
	)
}

type Message struct {
	Sender  string // required
	ReplyTo string // may be empty

	// At least one of these slices must have a non-zero length, except when
	// calling SendToAadmins.
	To, Cc, Bcc []string

	// At least one of Body or HTMLBody must by non-empty.
	Body     string
	HTMLBody string
}

// Send sends an email message.
func Send(msg *Message) error {
	err := smtp.SendMail(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		auth,
		msg.Sender,
		msg.To,
		[]byte(msg.Body),
	)

	if err != nil {
		return err
	}

	return nil
}
