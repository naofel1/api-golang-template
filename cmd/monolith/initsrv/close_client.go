package initsrv

import (
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// CloseClient will close all client connection
func (client *Client) CloseClient(ctx context.Context, logger *otelzap.Logger) {
	// Close the maria client when server stop
	if err := client.MariaDB.Close(); err != nil {
		logger.Ctx(ctx).Error("Failed to close mariadb database",
			zap.Error(err),
		)
	}
	// Close the redis client when server stop
	if err := client.Redis.Close(); err != nil {
		logger.Ctx(ctx).Error("Failed to close redis database",
			zap.Error(err),
		)
	}
	// Close the logger sync when server stop
	if err := logger.Sync(); err != nil {
		logger.Ctx(ctx).Error("Failed to close sync log",
			zap.Error(err),
		)
	}
	// Close the discord persistent connection when server stop
	if err := client.Discord.Close(); err != nil {
		logger.Ctx(ctx).Error("Failed to close discord WS Connection",
			zap.Error(err),
		)
	}
}
