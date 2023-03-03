package studentrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/ent/student"
)

// FindStudentByLogin return a Student found by his login
func (r *Repository) FindStudentByLogin(ctx context.Context, login string) (*ent.Student, error) {
	// Get the Student by login in the database
	StudentInfo, err := r.client.Student.
		Query().
		Where(student.Pseudo(login)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return StudentInfo, nil
}
