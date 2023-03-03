package studentservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"
	"github.com/naofel1/api-golang-template/pkg/pagination"

	"go.uber.org/zap"
)

// GetStudentsWithPagination retrieves all student at specific offset
func (s *studentService) GetStudentsWithPagination(ctx context.Context, u *ent.Students, pag *pagination.Front) error {
	enroll := func(ctx context.Context, repo Repository) error {
		studTotal, err := repo.GetTotalStudentCount(ctx)
		if err != nil {
			s.Logger.Ctx(ctx).Info("Error when getting students paginated",
				zap.Error(err),
			)

			return apistatus.NewInternal()
		}

		pag.Calculate(studTotal)

		if err := repo.GetAllStudentsAtOffset(ctx, u, pag.Offset, pag.ItemsPerPage); err != nil {
			s.Logger.Ctx(ctx).Info("Error when getting students paginated",
				zap.Error(err),
			)

			// Student is not found
			if ent.IsNotFound(err) {
				return apistatus.NewNotFound("student", "no one")
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
