package service

import (
	"context"
	"etl/internal/domain"
	"etl/internal/repository"
	"etl/internal/service/migrator"
	"github.com/pkg/errors"
)

type Service struct {
	migrator *migrator.Migrator
	repo     *repository.Repository
}

func NewService(m *migrator.Migrator, r *repository.Repository) *Service {
	return &Service{
		migrator: m,
		repo:     r,
	}
}

func (svc *Service) CreateDatabase(ctx context.Context, database *domain.Database) error {
	sql := database.ToSql()

	err := svc.repo.ExecSqlDest(ctx, sql)
	if err != nil {
		return errors.Wrap(err, "service.CreateDatabase")
	}
	return nil
}

func (svc *Service) MigrateData(ctx context.Context, routes *domain.Routes) error {
	if svc.migrator == nil {
		return errors.New("there is no migrator implementation")
	}
	err := svc.migrator.Migrate(ctx, routes)
	if err != nil {
		return errors.Wrap(err, "service.MigrateData")
	}
	return nil
}
