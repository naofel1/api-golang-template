package tracing

import (
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// SetOpenTelemetryInfo set all telemetry that will be collected
func SetOpenTelemetryInfo(ctx context.Context, logger *otelzap.Logger, tp trace.TracerProvider) {
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	logger.Ctx(ctx).Info("OpenTelemetry Provided set (Jaeger)")
}
