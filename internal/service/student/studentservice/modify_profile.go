package studentservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.uber.org/zap"
)

// ModifyProfile updates student information
func (s *studentService) ModifyProfile(ctx context.Context, u *ent.Student) error {
	enroll := func(ctx context.Context, repo Repository) error {
		if len(u.PasswordHash) != 0 {
			pw, err := generatePasswordHash(ctx, s.Tracer, string(u.PasswordHash))
			if err != nil {
				s.Logger.Ctx(ctx).Error("Unable to generate password for student",
					zap.Error(err),
				)

				return apistatus.NewInternal()
			}

			u.PasswordHash = pw
		}

		// Update the student with information specified
		if err := repo.UpdateStudent(ctx, u); err != nil {
			s.Logger.Ctx(ctx).Info("Error when Updating User",
				zap.String("Pseudo", u.Pseudo),
				zap.Error(err),
			)

			// Check if student cannot be updated as constraint error
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

	return nil
}
