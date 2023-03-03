// Package mailerrepo provides a repository to interact with the mailgun mailer
package mailerrepo

import (
	"github.com/naofel1/api-golang-template/internal/service/mailerservice"

	"github.com/mailgun/mailgun-go/v4"
)

// Repository is data/repository implementation
// of service layer TokenRepository
type Repository struct {
	MailerClient *mailgun.MailgunImpl
}

// New is a factory for initializing Token Repositories
func New(mailerClient *mailgun.MailgunImpl) *Repository {
	return &Repository{
		MailerClient: mailerClient,
	}
}

var _ mailerservice.Repository = (*Repository)(nil)
