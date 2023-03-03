package tracing

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/configs"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
)

// JaegerTraceProvider initialize endpoint where telemetry will be sent (Jaeger)
func JaegerTraceProvider(ctx context.Context, logger *otelzap.Logger, conf *configs.Config) *sdktrace.TracerProvider {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.Jaeger.ConnectionString)))
	if err != nil {
		logger.Ctx(ctx).Fatal("Jaeger error",
			zap.Error(err),
		)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("monolith"),
			semconv.DeploymentEnvironmentKey.String(conf.Host.Mode),
		)),
	)
}
