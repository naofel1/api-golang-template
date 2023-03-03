package studentrepo

import (
	"context"
)

// GetTotalStudentCount return count of total student
func (r *Repository) GetTotalStudentCount(ctx context.Context) (int, error) {
	totalStud, err := r.client.Student.
		Query().
		Count(ctx)
	if err != nil {
		return 0, err
	}

	return totalStud, nil
}
