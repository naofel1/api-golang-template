package studentrepo

import (
	"context"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/ent/student"
)

// GetTotalStudentByPseudoOrFirstNameOrLastNameCount return count of total student
func (r *Repository) GetTotalStudentByPseudoOrFirstNameOrLastNameCount(ctx context.Context, n *ent.Student) (int, error) {
	totalStud, err := r.client.Student.
		Query().
		Where(student.Or(
			student.PseudoContainsFold(n.Pseudo),
			student.FirstNameContainsFold(n.FirstName),
			student.LastNameContainsFold(n.LastName),
		)).
		Count(ctx)
	if err != nil {
		return 0, err
	}

	return totalStud, nil
}
