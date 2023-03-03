package bucketrepo

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadFile upload a new file in S3 bucket
func (r *Repository) UploadFile(ctx context.Context, bucketName, objName string, imageFile multipart.File) (string, error) {
	_, err := r.S3Client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("healthcheck/" + objName),
		Body:   imageFile,
	})
	if err != nil {
		return "", err
	}

	return "ok", nil
}
