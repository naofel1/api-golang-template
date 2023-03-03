package adminrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/ent/admin"
)

// FindAdminByLogin return an Admin found by his login.
func (r *Repository) FindAdminByLogin(ctx context.Context, login string) (*ent.Admin, error) {
	// Get the Admin by login in the database
	AdminInfo, err := r.client.Admin.
		Query().
		Where(admin.Pseudo(login)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return AdminInfo, nil
}
