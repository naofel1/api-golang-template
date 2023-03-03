package adminrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/ent/admin"

	"github.com/google/uuid"
)

// FindAdminByID return an Admin found by his UUID.
func (r *Repository) FindAdminByID(ctx context.Context, uid uuid.UUID) (*ent.Admin, error) {
	// Check if ID exists
	AdminInfo, err := r.client.Admin.
		Query().
		Where(admin.ID(uid)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return AdminInfo, nil
}
