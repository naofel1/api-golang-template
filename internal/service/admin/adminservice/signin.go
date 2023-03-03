package adminservice

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/pkg/apistatus"

	"go.uber.org/zap"
)

// Signin reaches our AdminRepository check if the admin exists
// and then compares the supplied password with the provided password
// if a valid email/password combo is provided, u will hold all
// available user fields
func (s *adminService) Signin(ctx context.Context, u *ent.Admin) error {
	ctx, span := s.Tracer.Start(ctx, "signin")
	defer span.End()

	enroll := func(ctx context.Context, repo Repository) error {
		uFetched, err := repo.FindAdminByLogin(ctx, u.Pseudo)
		// Will return NotAuthorized to client to omit details of why
		if err != nil {
			s.Logger.Ctx(ctx).Info("Error when finding User by Email",
				zap.String("Pseudo", u.Pseudo),
				zap.Error(err),
			)
			// Admin is not found
			if ent.IsNotFound(err) {
				return apistatus.NewAuthorization("Invalid email and password combination")
			}

			return apistatus.NewInternal()
		}

		// verify password - we previously created this method
		match, err := validatePasswordHash(uFetched.PasswordHash, u.PasswordHash)
		if err != nil {
			s.Logger.Ctx(ctx).Info("Cannot validate password hash",
				zap.Error(err),
			)

			return apistatus.NewInternal()
		}
		// Check if the password given match to the password stored
		if !match {
			return apistatus.NewAuthorization("Invalid email and password combination")
		}

		*u = *uFetched

		return nil
	}

	if err := s.AdminRepository.Execute(ctx, enroll); err != nil {
		return err
	}

	return nil
}
