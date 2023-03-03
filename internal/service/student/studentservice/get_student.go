package studentservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.uber.org/zap"
)

// GetStudent retrieves a student based on their uuid
func (s *studentService) GetStudent(ctx context.Context, u *ent.Student) error {
	enroll := func(ctx context.Context, repo Repository) error {
		uFetched, err := repo.FindStudentByID(ctx, u.ID)
		if err != nil {
			s.Logger.Ctx(ctx).Info("Error when getting User",
				zap.String("UUID", u.ID.String()),
				zap.Error(err),
			)

			// Student is not found
			if ent.IsNotFound(err) {
				return apistatus.NewNotFound("UUID", u.ID.String())
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
