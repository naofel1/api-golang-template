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

			/*
				Angus: Error handling is such a huge topic! This isn't necessarily a bad way of handling errors, but be
				aware that you're very close to mixing domain layer and transport layer concerns.

				I would argue that the domain has no concept of a generic "internal error". The domain knows exactly
				what's gone wrong! It also doesn't know that it's returning its value to an HTTP handler. It might be
				going to gRPC, or a CLI, or a message queue, etc.

				The cleanest thing to do when you receive an error you weren't expecting, it is pass that error up the
				chain to the handler. If the handler doesn't know what to do with it, pass the error up to the global
				recovery middleware (either by returning the error or panicking, depending on the web framework you're
				using). The recovery middleware is the HTTP adapter concern that will handle the translation of an
				unknown error to a 500 Internal Server Error. Your HTTP server logging middleware will also eliminate
				the need to log in this function.

				My preferred approach to error handling is to define domain error types for the kinds of errors that you
				expect to handle. E.g. AuthError, EntityNotFound, EntityAlreadyExists, etc. The repository layer is
				responsible for mapping its own errors into these domain error types. If you receive an error in a
				repository (or from some other third-party package) that DOESN'T map to a domain error types, it gets
				passed all the way back up the chain to the global recovery handler, which will translate it to a 500.

			*/
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
