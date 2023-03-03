package mailerrepo

import (
	"context"

	"github.com/mailgun/mailgun-go/v4"
)

// SendMail will send an email with mailgun client
func (r *Repository) SendMail(ctx context.Context, messages *mailgun.Message) (string, error) {
	// If the user is not present, we will add the user and also associate the log
	_, id, err := r.MailerClient.Send(ctx, messages)
	if err != nil {
		return "", err
	}

	return id, nil
}
