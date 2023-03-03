package studentservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.uber.org/zap"
)

// DeleteStudent will definitively delete in the db the student entity
func (s *studentService) DeleteStudent(ctx context.Context, u *ent.Student) error {
	enroll := func(ctx context.Context, repo Repository) error {
		// Update the student with information specified
		if err := repo.DeleteStudent(ctx, u); err != nil {
			s.Logger.Ctx(ctx).Info("Error when deleting Student",
				zap.String("Pseudo", u.Pseudo),
				zap.Error(err),
			)

			return apistatus.NewInternal()
		}

		return nil
	}

	if err := s.StudentRepository.Execute(ctx, enroll); err != nil {
		return err
	}

	return nil
}
