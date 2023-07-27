package bucketservice

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"go.uber.org/zap"
)

// UploadHealthCheckResult will upload on a S3 Bucket, the result of health check (connectivity)
func (s *bucketService) UploadHealthCheckResult(ctx context.Context, bucketName, objName string, imageFileHeader *multipart.FileHeader) (string, error) {
	ctx, span := s.Tracer.Start(ctx, "Upload HealthCheck Result")
	defer span.End()

	imageFile, err := imageFileHeader.Open()
	if err != nil {
		s.Logger.Ctx(ctx).Info("Failed to open image file", zap.Error(err))
		return "", apistatus.NewInternal()
	}

	link, err := s.BucketRepository.UploadFile(ctx, bucketName, objName, imageFile)
	if err != nil {
		s.Logger.Error("Cannot upload file", zap.Error(err))

		/*
			Angus: Careful! You have a repository error creeping into your domain layer. The domain shouldn't know
			anything about AWS or S3 :) The BucketRepository implementation should translate its errors into error types
			defined by the domain.

			Naofel: Oh, wow! I hadn't realized that. I'll make sure to correct this oversight about the repository error
			seeping into the domain layer in my upcoming release.
		*/
		var awsErr awserr.Error
		if ok := errors.As(err, &awsErr); ok && awsErr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			s.Logger.Ctx(ctx).Info("upload canceled due to timeout", zap.Error(err))
		} else {
			s.Logger.Ctx(ctx).Info("failed to upload object", zap.Error(err))
		}

		return "", apistatus.NewInternal()
	}

	return link, nil
}
