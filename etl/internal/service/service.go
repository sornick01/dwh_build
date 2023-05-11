package service

import (
	"etl/internal/repository"
	"etl/internal/service/migrator"
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
