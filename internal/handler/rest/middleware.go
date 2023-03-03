package rest

import (
	"net/http"
	"time"

	"github.com/naofel1/api-golang-template/pkg/slice"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ServerMiddleware represents a chain of middleware in the order in which
// they'll be applied to a *gin.Engine. I.e. the first middleware in the stack
// represents the outermost layer in the HTTP call chain.
type ServerMiddleware []gin.HandlerFunc

// GlobalServerMiddleware returns the middleware stack used for all routes
// when running in development.
func GlobalServerMiddleware(logger *otelzap.Logger) ServerMiddleware {
	return ServerMiddleware{
		otelgin.Middleware("Global Server Middleware"),
		ginZapLogger(logger),
		ginzap.RecoveryWithZap(logger.Logger, true),
		contentTypes(applicationJSON),
	}
}

type contentType string

const (
	applicationJSON contentType = "application/json"
)

func contentTypes(contentTypes ...contentType) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.ContentType() == "" {
			return
		}

		if !slice.Includes(contentTypes, contentType(c.ContentType())) {
			c.AbortWithStatus(http.StatusUnsupportedMediaType)

			return
		}

		c.Next()
	}
}

func ginZapLogger(logger *otelzap.Logger) gin.HandlerFunc {
	return ginzap.GinzapWithConfig(logger.Logger, &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
		Context: func(c *gin.Context) []zapcore.Field {
			var fields []zapcore.Field
			// log request ID
			if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
				fields = append(fields, zap.String("request_id", requestID))
			}

			// log trace and span ID
			if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
				fields = append(fields,
					zap.String("trace_id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()),
					zap.String("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()),
				)
			}

			return fields
		},
	})
}
