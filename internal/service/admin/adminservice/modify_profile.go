package adminservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.uber.org/zap"
)

// ModifyProfile updates admin information in database
func (s *adminService) ModifyProfile(ctx context.Context, u *ent.Admin) error {
	enroll := func(ctx context.Context, repo Repository) error {
		if len(u.PasswordHash) != 0 {
			pw, err := generatePasswordHash(string(u.PasswordHash))
			if err != nil {
				s.Logger.Ctx(ctx).Error("Unable to generate password for admin",
					zap.Error(err),
				)

				return apistatus.NewInternal()
			}

			u.PasswordHash = pw
		}
		// Update admin information
		if err := repo.UpdateAdmin(ctx, u); err != nil {
			s.Logger.Ctx(ctx).Info("Error when Updating User",
				zap.String("Admin UUID", u.ID.String()),
				zap.Error(err),
			)

			return apistatus.NewInternal()
		}

		return nil
	}
	if err := s.AdminRepository.Execute(ctx, enroll); err != nil {
		return err
	}

	return nil
}
