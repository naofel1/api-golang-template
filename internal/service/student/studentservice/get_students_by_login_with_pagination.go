package studentservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"
	"github.com/naofel1/api-golang-template/pkg/pagination"

	"go.uber.org/zap"
)

// GetStudentsWithSearchWithPagination retrieves a student based on their pseudo
func (s *studentService) GetStudentsWithSearchWithPagination(ctx context.Context, u *ent.Students, us *ent.Student, pag *pagination.Front) error {
	enroll := func(ctx context.Context, repo Repository) error {
		studTotal, err := repo.GetTotalStudentByPseudoOrFirstNameOrLastNameCount(ctx, us)
		if err != nil {
			s.Logger.Ctx(ctx).Info("Error when getting students paginated",
				zap.Error(err),
			)

			return apistatus.NewInternal()
		}

		pag.Calculate(studTotal)

		uFetched, err := repo.FindStudentsByLoginOrFirstNameOrLastNameAtOffset(ctx, us, pag.Offset, pag.ItemsPerPage)
		if err != nil {
			s.Logger.Ctx(ctx).Info("Error when getting User",
				zap.String("Pseudo", us.Pseudo),
				zap.Error(err),
			)

			// Student is not found
			if ent.IsNotFound(err) {
				return apistatus.NewNotFound("Pseudo", us.Pseudo)
			}

			return apistatus.NewInternal()
		}

		*u = uFetched

		return nil
	}
	if err := s.StudentRepository.Execute(ctx, enroll); err != nil {
		return err
	}

	return nil
}
