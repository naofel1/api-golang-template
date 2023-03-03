package studentservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// Signup reaches our StudentRepository to verify the
// pseudo is available and signs up the user if this is the case
func (s *studentService) Signup(ctx context.Context, u *ent.Student) error {
	ctx, span := s.Tracer.Start(ctx, "signup")
	defer span.End()

	if len(u.PasswordHash) != 0 {
		pw, err := generatePasswordHash(ctx, s.Tracer, string(u.PasswordHash))
		if err != nil {
			s.Logger.Ctx(ctx).Error("Unable to signup student profile",
				zap.String("Pseudo", u.Pseudo),
			)

			return apistatus.NewInternal()
		}

		// now I realize why I originally used Signup(ctx, email, password)
		// then created a user. It's somewhat un-natural to mutate the user here
		u.PasswordHash = pw
	}

	enroll := func(ctx context.Context, repo Repository) error {
		if err := repo.CreateStudent(ctx, u); err != nil {
			s.Logger.Ctx(ctx).Info("Error when Register User",
				zap.String("Pseudo", u.Pseudo),
				zap.Error(err),
			)

			if ent.IsConstraintError(err) {
				return apistatus.NewConflict("pseudo", u.Pseudo)
			}

			return apistatus.NewInternal()
		}

		return nil
	}

	if err := s.StudentRepository.Execute(ctx, enroll); err != nil {
		return err
	}

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("enduser.id", u.ID.String()),
			attribute.String("enduser.pseudo", u.Pseudo),
		)
	}

	return nil
}
