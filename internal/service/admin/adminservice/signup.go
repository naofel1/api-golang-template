package adminservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.uber.org/zap"
)

// Signup reaches our AdminRepository to verify the
// pseudo is available and signs up the user if this is the case
func (s *adminService) Signup(ctx context.Context, u *ent.Admin) error {
	ctx, span := s.Tracer.Start(ctx, "signup")
	defer span.End()

	pw, err := generatePasswordHash(string(u.PasswordHash))
	if err != nil {
		s.Logger.Ctx(ctx).Error("Unable to signup admin profile",
			zap.String("Pseudo", u.Pseudo),
		)

		return apistatus.NewInternal()
	}

	// now I realize why I originally used Signup(ctx, email, password)
	// then created a user. It's somewhat un-natural to mutate the user here
	u.PasswordHash = pw

	enroll := func(ctx context.Context, repo Repository) error {
		if err := repo.CreateAdmin(ctx, u); err != nil {
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

	if err := s.AdminRepository.Execute(ctx, enroll); err != nil {
		return err
	}

	return nil
}
