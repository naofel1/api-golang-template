package studentrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/ent/student"

	"github.com/google/uuid"
)

// FindStudentByID return a Student found by his UUID
func (r *Repository) FindStudentByID(ctx context.Context, uid uuid.UUID) (*ent.Student, error) {
	// Check if ID exists
	StudentInfo, err := r.client.Student.
		Query().
		Where(student.ID(uid)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return StudentInfo, nil
}
