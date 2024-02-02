package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool `json:"-"`
}

func NewRepository(p *pgxpool.Pool) *Repository {
	return &Repository{pool: p}
}

func (r *Repository) ExecSql(ctx context.Context, sql string) error {
	_, err := r.pool.Exec(ctx, sql)
	if err != nil {
		return fmt.Errorf("repository.ExecSql(): %w", err)
	}
	return nil
}
