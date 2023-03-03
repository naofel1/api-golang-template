package studentrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
)

// CreateStudent create in the database a Student with the provided information
func (r *Repository) CreateStudent(ctx context.Context, u *ent.Student) error {
	// Create a new student entity
	createdStudent, err := r.client.Student.
		Create().
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetPseudo(u.Pseudo).
		SetGender(u.Gender).
		SetPasswordHash(u.PasswordHash).
		Save(ctx)
	if err != nil {
		return err
	}

	*u = *createdStudent

	return nil
}
