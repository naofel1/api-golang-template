package studentrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/ent/student"
)

// CheckStudentPseudo check in the database if the student pseudo is available
func (r *Repository) CheckStudentPseudo(ctx context.Context, u *ent.Student) (bool, error) {
	// Create a new student entity
	ok, err := r.client.Student.
		Query().
		Where(student.PseudoEqualFold(u.Pseudo)).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return ok, nil
}
