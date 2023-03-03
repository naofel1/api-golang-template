package studentrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/ent/student"
)

// GetAllStudentsAtOffset return all the student at the specified offset
func (r *Repository) GetAllStudentsAtOffset(ctx context.Context, u *ent.Students, off, limit int) error {
	sFetched, err := r.client.Student.
		Query().
		Order(ent.Desc(student.FieldCreatedAt)).
		Offset(off).
		Limit(limit).
		All(ctx)
	if err != nil {
		return err
	}

	*u = sFetched

	return nil
}
