package studentrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
)

// DeleteStudent delete in the database a Student
func (r *Repository) DeleteStudent(ctx context.Context, u *ent.Student) error {
	// Create a new student entity
	err := r.client.Student.
		DeleteOne(u).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
