package studentrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
)

// UpdateStudent update in the database a Student with the provided information
func (r *Repository) UpdateStudent(ctx context.Context, u *ent.Student) error {
	sUpdated, err := r.client.Student.
		UpdateOne(u).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetPseudo(u.Pseudo).
		SetGender(u.Gender).
		SetBirthday(u.Birthday).
		SetPasswordHash(u.PasswordHash).
		Save(ctx)
	if err != nil {
		return err
	}

	*u = *sUpdated

	return nil
}
