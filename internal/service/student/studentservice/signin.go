package studentservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// Signin reaches our StudentRepository check if the student exists
// and then compares the supplied password with the provided password
// if a valid email/password combo is provided, u will hold all
// available user fields
func (s *studentService) Signin(ctx context.Context, u *ent.Student) error {
	ctx, span := s.Tracer.Start(ctx, "Signin")
	defer span.End()

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("enduser.pseudo", u.Pseudo),
		)
	}

	enroll := func(ctx context.Context, repo Repository) error {
		uFetched, err := repo.FindStudentByLogin(ctx, u.Pseudo)
		// Will return NotAuthorized to client to omit details of why
		if err != nil {
			s.Logger.Ctx(ctx).Info("Error when finding User by Email",
				zap.String("Pseudo", u.Pseudo),
				zap.Error(err),
			)

			if ent.IsNotFound(err) {
				return apistatus.NewAuthorization("Invalid email and password combination")
			}

			return apistatus.NewInternal()
		}

		if uFetched.PasswordHash == nil {
			s.Logger.Ctx(ctx).Info("Password not set for student")

			return apistatus.NewAuthorization("Invalid email and password combination")
		}

		// verify password - we previously created this method
		match, err := validatePasswordHash(ctx, s.Tracer, uFetched.PasswordHash, u.PasswordHash)
		if err != nil {
			s.Logger.Ctx(ctx).Info("Cannot validate password hash",
				zap.Error(err),
			)

			return apistatus.NewInternal()
		}
		// Check if password match
		if !match {
			return apistatus.NewAuthorization("Invalid email and password combination")
		}

		*u = *uFetched

		return nil
	}

	if err := s.StudentRepository.Execute(ctx, enroll); err != nil {
		return err
	}

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("enduser.id", u.ID.String()),
		)
	}

	return nil
}
