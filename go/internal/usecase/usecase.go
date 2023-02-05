package usecase

import (
	"dwh/internal/usecase/creation"
	"dwh/internal/usecase/etl"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UseCase interface {
}

type UseCaseStruct struct {
	*creation.Creator
	*etl.ETL
}

func New(p *pgxpool.Pool) UseCase {
	return &UseCaseStruct{
		Creator: creation.New(p),
		ETL:     etl.New(p),
	}
}
