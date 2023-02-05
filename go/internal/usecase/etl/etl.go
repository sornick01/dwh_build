package etl

import "github.com/jackc/pgx/v5/pgxpool"

type ETL struct {
	pool *pgxpool.Pool
}

func New(p *pgxpool.Pool) *ETL {
	return &ETL{pool: p}
}
