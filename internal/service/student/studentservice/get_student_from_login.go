package studentservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.uber.org/zap"
)

// GetStudentFromLogin retrieves a student based on their pseudo
func (s *studentService) GetStudentFromLogin(ctx context.Context, u *ent.Student) error {
	enroll := func(ctx context.Context, repo Repository) error {
		uFetched, err := repo.FindStudentByLogin(ctx, u.Pseudo)
		if err != nil {
			s.Logger.Ctx(ctx).Info("Error when getting User",
				zap.String("Pseudo", u.Pseudo),
				zap.Error(err),
			)

			// Student is not found
			if ent.IsNotFound(err) {
				return apistatus.NewNotFound("Pseudo", u.Pseudo)
			}

			return apistatus.NewInternal()
		}

		*u = *uFetched

		return nil
	}
	if err := s.StudentRepository.Execute(ctx, enroll); err != nil {
		return err
	}

	return nil
}
