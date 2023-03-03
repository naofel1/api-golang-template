package initsrv

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/configs"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

// UtilsConfig will hold utils that will be injected
// into Service layer on service initialization
type UtilsConfig struct {
	Certs *configs.JWTCertificate
}

// InitUtils will initialize all utils needed by application
func InitUtils(ctx context.Context, logger *otelzap.Logger, conf *configs.Config) *UtilsConfig {
	return &UtilsConfig{
		// Load RSA key used for generate JWT Token
		Certs: configs.NewRSACerts(ctx, logger, conf.Certs),
	}
}
