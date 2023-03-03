package adminrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
)

// CreateAdmin create in the database an Admin with the provided information.
func (r *Repository) CreateAdmin(ctx context.Context, u *ent.Admin) error {
	createdAdmin, err := r.client.Admin.
		Create().
		SetPseudo(u.Pseudo).
		SetPasswordHash(u.PasswordHash).
		Save(ctx)
	if err != nil {
		return err
	}

	*u = *createdAdmin

	return nil
}
