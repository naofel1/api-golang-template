package mailerrepo

import (
	"github.com/mailgun/mailgun-go/v4"
)

// NewMessage will return a new message with mailgun client
func (r *Repository) NewMessage(sender, subject, recipient string) *mailgun.Message {
	return r.MailerClient.NewMessage(sender, subject, "", recipient)
}
