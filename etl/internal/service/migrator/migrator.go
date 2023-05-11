package migrator

import (
	"context"
	"errors"
	"etl/internal/config"
	"etl/internal/domain"
	"etl/internal/repository"
	"etl/internal/service/validator"
	"etl/pkg/ptrconv"
	"strings"
)

type Migrator struct {
	validator *validator.Validator
	//routes    *domain.Routes
	repo *repository.Repository
}

func NewMigrator(val *validator.Validator, repo *repository.Repository) *Migrator {
	return &Migrator{
		validator: val,
		repo:      repo,
	}
}

// Migrate - запуск миграции данных
func (m *Migrator) Migrate(ctx context.Context, routes *domain.Routes) error {
	// TODO: сделать получение кол-ва строк из исходной таблицы
	if routes == nil {
		return errors.New("empty routes")
	}
	rowsCount, err := m.repo.TotalSourceRows(ctx, routes.SourceSQL, routes.Filters)
	if err != nil {
		return err
	}

	if routes.Filters == nil {
		routes.Filters = &domain.Filters{}
	}
	if routes.Filters.Limit == nil {
		routes.Filters.Limit = ptrconv.Int(config.BatchSize)
	}
	if routes.Filters.Offset == nil {
		routes.Filters.Offset = ptrconv.Int(0)
	}
	// миграция данных батчами
	for i := 0; i < rowsCount; i += config.BatchSize {
		// TODO: добавить тело цикла в транзакцию ???
		//select
		var rows []map[string]interface{}
		srcColumnsList := routes.ListSrcColumns()
		srcColumns := strings.Join(srcColumnsList, ",")
		rows, err = m.repo.GetSourceRows(ctx, routes.Filters, routes.SourceSQL, srcColumns)
		if err != nil {
			return err
		}

		// validate
		rows, err = m.validator.Run(ctx, rows)
		if err != nil {
			return err
		}

		// insert
		routeMap := routes.GetRouteMap()
		targetColumns := routes.ListTargetColumns()
		err = m.repo.InsertRows(ctx, rows, routeMap, targetColumns, routes.TargetTable.String())
		if err != nil {
			return err
		}
		routes.Filters.Offset = ptrconv.Int(*routes.Filters.Offset + config.BatchSize)
	}
	return nil
}
