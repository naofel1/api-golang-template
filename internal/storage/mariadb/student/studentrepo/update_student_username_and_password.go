package studentrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
)

// UpdatePseudoUsernameAndPassword update in the database a Student with the provided information
func (r *Repository) UpdatePseudoUsernameAndPassword(ctx context.Context, u *ent.Student) error {
	sUpdated, err := r.client.Student.
		UpdateOne(u).
		SetPseudo(u.Pseudo).
		SetPasswordHash(u.PasswordHash).
		Save(ctx)
	if err != nil {
		return err
	}

	*u = *sUpdated

	return nil
}
