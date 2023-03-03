package client

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/configs"

	"github.com/centrifugal/gocent/v3"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// InitCentrifugo will return a centrifugo client initialized (PubSub)
func InitCentrifugo(ctx context.Context, logger *otelzap.Logger, conf *configs.Centrifugo) *gocent.Client {
	client := gocent.New(gocent.Config{
		Addr: conf.ConnectionString,
		Key:  conf.APIKey,
	})

	logger.Ctx(ctx).Info("Centrifugo client initialized",
		zap.String("Connection info", conf.ConnectionString),
	)

	return client
}
