package adminservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.uber.org/zap"
)

// GetAdmin retrieves an admin based on their uuid.
func (s *adminService) GetAdmin(ctx context.Context, u *ent.Admin) error {
	enroll := func(ctx context.Context, repo Repository) error {
		uFetched, err := repo.FindAdminByID(ctx, u.ID)
		if err != nil {
			s.Logger.Ctx(ctx).Info("Error when getting User",
				zap.String("UUID", u.ID.String()),
				zap.Error(err),
			)

			if ent.IsNotFound(err) {
				return apistatus.NewNotFound("UUID", u.ID.String())
			}

			return apistatus.NewInternal()
		}

		*u = *uFetched

		return nil
	}
	if err := s.AdminRepository.Execute(ctx, enroll); err != nil {
		return err
	}

	return nil
}
