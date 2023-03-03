package studentrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/ent/student"
)

// FindStudentsByLoginOrFirstNameOrLastNameAtOffset return al Studenta found by his login
func (r *Repository) FindStudentsByLoginOrFirstNameOrLastNameAtOffset(ctx context.Context, stud *ent.Student, off, limit int) (ent.Students, error) {
	// Get the Student by login in the database
	fetchedStudent, err := r.client.Student.
		Query().
		Where(student.Or(
			student.FirstNameContainsFold(stud.FirstName),
			student.LastNameContainsFold(stud.LastName),
			student.PseudoContainsFold(stud.Pseudo),
		)).
		Order(ent.Desc(student.FieldCreatedAt)).
		Offset(off).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return fetchedStudent, nil
}
