package repo

import "github.com/jackc/pgx/v5/pgxpool"

type Repo struct {
	Pool *pgxpool.Pool
}
