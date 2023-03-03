package adminrepo

import (
	"context"
	"fmt"

	"github.com/naofel1/api-golang-template/internal/ent"
	"github.com/naofel1/api-golang-template/internal/service/admin/adminservice"
)

// AtomicRepository satisfies adminservice.AtomicRepository.
type AtomicRepository struct {
	db *ent.Client
}

var _ adminservice.AtomicRepository = (*AtomicRepository)(nil)

// NewAtomic instantiates a new AtomicRepository using the database provided.
func NewAtomic(db *ent.Client) *AtomicRepository {
	return &AtomicRepository{db: db}
}

// Execute decorates the given AtomicOperation with a transaction. If the
// AtomicOperation returns an error, the transaction is rolled back. Otherwise,
// the transaction is committed.
func (ar *AtomicRepository) Execute(
	ctx context.Context,
	op adminservice.AtomicOperation,
) error {
	tx, err := ar.db.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()

			panic(v)
		}
	}()

	adminRepoWithTransaction := Repository{client: tx}

	if err := op(ctx, &adminRepoWithTransaction); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

// Repository satisfies adminservice.Repository. It is agnostic whether
// its sql.TableOperator is a database or transaction.
type Repository struct {
	client *ent.Tx
}

var _ adminservice.Repository = (*Repository)(nil)
