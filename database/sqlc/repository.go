package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type DatabaseRepository struct {
	*Queries
	Db *pgxpool.Pool
}

func NewDatabaseRepository(Db *pgxpool.Pool) *DatabaseRepository {
	return &DatabaseRepository{
		Queries: New(Db),
		Db:      Db,
	}
}

func (r *DatabaseRepository) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := r.Db.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit(ctx)
}
