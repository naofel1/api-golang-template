package client

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/configs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

// InitAWS return a new instance of an AWS Client
func InitAWS(ctx context.Context, logger *otelzap.Logger, cfg *configs.AWS) *s3.S3 {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.Must(session.NewSession())

	// here, we provide access and secret keys for aws
	creds := credentials.NewStaticCredentials(cfg.SCWAccessKey, cfg.SCWSecretKey, "")

	s3Client := s3.New(sess, aws.NewConfig().WithRegion(cfg.Region).WithEndpoint(cfg.Endpoint).WithCredentials(creds))

	logger.Ctx(ctx).Info("AWS client initialized...")

	return s3Client
}
