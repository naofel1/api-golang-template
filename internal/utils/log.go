package utils

import (
	"context"
	"log"

	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/naofel1/api-golang-template/internal/primitive"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger return a new instance of a Zap Logger
func InitLogger(ctx context.Context, cfg *configs.AppInfo) *otelzap.Logger {
	var cfgLogs zap.Config

	// Choose between dev and prod logs
	if cfg.Mode == primitive.AppModeDev.String() {
		cfgLogs = zap.NewDevelopmentConfig()
		cfgLogs.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfgLogs = zap.NewProductionConfig()
	}

	// Build the log config
	logger, err := cfgLogs.Build()
	if err != nil {
		log.Fatalf("failed init Zap Log: %v", err)
	}

	// Wraper for OpenTelemetry
	loggerWrapped := otelzap.New(logger,
		otelzap.WithMinLevel(zap.InfoLevel),
		otelzap.WithTraceIDField(true))

	loggerWrapped.Ctx(ctx).Info("Logger initialized")

	return loggerWrapped
}
