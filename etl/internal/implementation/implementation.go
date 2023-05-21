package implementation

import (
	"context"
	"encoding/json"
	"etl/etl/api/etl"
	"etl/internal/domain"
	"etl/internal/repository"
	"etl/internal/service"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Implementation struct {
	svc  *service.Service
	repo *repository.Repository
	etl.UnimplementedEtlServer
}

func NewImplementation(svc *service.Service, repo *repository.Repository) Implementation {
	return Implementation{
		svc:  svc,
		repo: repo,
	}
}

func (impl Implementation) CreateDatabase(ctx context.Context, req *etl.CreateDatabaseRequest) (*etl.CreateDatabaseResponse, error) {
	idCfg, err := uuid.Parse(req.ConfigId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var cfg string
	cfg, err = impl.repo.GetConfig(ctx, idCfg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var database domain.Database
	err = json.Unmarshal([]byte(cfg), &database)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = impl.svc.CreateDatabase(ctx, &database)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &etl.CreateDatabaseResponse{Success: true}, nil
}

func (impl Implementation) MigrateData(ctx context.Context, req *etl.MigrateDataRequest) (*etl.MigrateDataResponse, error) {
	idCfg, err := uuid.Parse(req.ConfigId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var cfg string
	cfg, err = impl.repo.GetConfig(ctx, idCfg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	routes := &domain.Routes{}
	err = json.Unmarshal([]byte(cfg), routes)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = impl.svc.MigrateData(ctx, routes)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &etl.MigrateDataResponse{}, nil
}
