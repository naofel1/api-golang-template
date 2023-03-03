package bucketservice

import (
	"context"
	"mime/multipart"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
)

// Interface specifies the business operations of the service.
type Interface interface {
	UploadHealthCheckResult(ctx context.Context, bucketName, objName string, imageFileHeader *multipart.FileHeader) (string, error)
}

// Config will hold repository and used utils that will be injected
// into this Service layer on service initialization
type Config struct {
	BucketRepository Repository
	Tracer           trace.Tracer
	Logger           *otelzap.Logger
}

// New configures and returns an Interface implementation.
func New(c *Config) Interface {
	return &bucketService{
		BucketRepository: c.BucketRepository,
		Tracer:           c.Tracer,
		Logger:           c.Logger,
	}
}

// tokenService implements tokenService.Interface.
type bucketService struct {
	BucketRepository Repository
	Tracer           trace.Tracer
	Logger           *otelzap.Logger
}

// Repository contain all the function available in the defined domain
type Repository interface {
	UploadFile(ctx context.Context, bucketName, objName string, imageFile multipart.File) (string, error)
}
