package adminrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
)

// UpdateAdmin update in the database an Admin with the provided information
func (r *Repository) UpdateAdmin(ctx context.Context, u *ent.Admin) error {
	tUpdated, err := r.client.Admin.
		UpdateOne(u).
		SetPasswordHash(u.PasswordHash).
		Save(ctx)
	if err != nil {
		return err
	}

	*u = *tUpdated

	return nil
}
