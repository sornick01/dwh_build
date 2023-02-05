package creation

import "github.com/jackc/pgx/v5/pgxpool"

type Creator struct {
}

func New(p *pgxpool.Pool) *Creator {
	return &Creator{pool: p}
}
