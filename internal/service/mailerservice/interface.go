package mailerservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/configs"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
)

// Interface specifies the business operations of the service.
type Interface interface{}

// Config will hold repository and used utils that will be injected
// into this Service layer on service initialization
type Config struct {
	MailerRepo Repository

	MailConfig *configs.MailgunMailer
	Tracer     trace.Tracer
	Logger     *otelzap.Logger
	Conf       *configs.Host
}

// New configures and returns an Interface implementation.
func New(c *Config) Interface {
	return &mailerService{
		MailerRepository: c.MailerRepo,
		Tracer:           c.Tracer,
		Conf:             c.Conf,
		Logger:           c.Logger,
	}
}

// mailerService implements mailerService.Interface.
type mailerService struct {
	MailerRepository Repository
	Tracer           trace.Tracer
	Logger           *otelzap.Logger
	Conf             *configs.Host
}

// Repository contain all the function available in the defined domain
type Repository interface {
	NewMessage(sender, subject, recipient string) *mailgun.Message
	SendMail(ctx context.Context, messages *mailgun.Message) (string, error)
}
