package bucketrepo

import (
	"github.com/naofel1/api-golang-template/internal/service/bucket/bucketservice"

	"github.com/aws/aws-sdk-go/service/s3"
)

// Repository is data/repository implementation
// of service layer TokenRepository
type Repository struct {
	S3Client *s3.S3
}

// New is a factory for initializing Token Repositories
func New(s3Client *s3.S3) *Repository {
	return &Repository{
		S3Client: s3Client,
	}
}

var _ bucketservice.Repository = (*Repository)(nil)
