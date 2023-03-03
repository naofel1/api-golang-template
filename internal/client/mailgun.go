package client

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/configs"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// InitMailGun will return a mailgun client initialized
func InitMailGun(ctx context.Context, logger *otelzap.Logger, conf *configs.Mailgun) *mailgun.MailgunImpl {
	mg := mailgun.NewMailgun(conf.ClientDomain, conf.ClientSecret)
	mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	logger.Ctx(ctx).Info("Successfully init mailgun client", zap.String("Domain", conf.ClientDomain))

	return mg
}
