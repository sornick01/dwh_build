package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"

	"etl/internal/domain"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Tx pgx.Tx

type Repository struct {
	cfgs *pgxpool.Pool // База с конфигами
	src  *pgxpool.Pool // База-источник
	dest *pgxpool.Pool // Целевая база
}

func NewRepository(cfgs, src, dest *pgxpool.Pool) *Repository {
	return &Repository{
		cfgs: cfgs,
		src:  src,
		dest: dest,
	}
}

// ExecSqlDest - выполнить sql скрипт в целевой базе
func (r *Repository) ExecSqlDest(ctx context.Context, sql string) error {
	_, err := r.dest.Exec(ctx, sql)
	if err != nil {
		return fmt.Errorf("repository.ExecSqlDest(): %w", err)
	}
	return nil
}

// TotalSourceRows - кол-во строк-источников
func (r *Repository) TotalSourceRows(ctx context.Context, src string, filters *domain.Filters) (int, error) {
	query := psql().Select("COUNT(*)").From(src)

	query = r.applyFilters(ctx, query, filters)

	stmt, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	row := r.src.QueryRow(ctx, stmt, args)
	var n int
	err = row.Scan(&n)
	if err != nil {
		return 0, err
	}

	return n, nil
}

// GetConfig - получение конфига по id
func (r *Repository) GetConfig(ctx context.Context, id uuid.UUID) (string, error) {
	query := psql().Select("config").From("configs").Where(sq.Eq{"id": id})

	stmt, args, err := query.ToSql()
	if err != nil {
		return "", err
	}

	row := r.cfgs.QueryRow(ctx, stmt, args)
	var cfg string
	err = row.Scan(&cfg)
	if err != nil {
		return "", err
	}

	return cfg, nil
}

// GetSourceRows - получение строк-источников
func (r *Repository) GetSourceRows(ctx context.Context, filters *domain.Filters, src, srcColumns string) ([]map[string]interface{}, error) {
	query := psql().Select(srcColumns).From(src)

	query = r.applyFilters(ctx, query, filters)

	stmt, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.src.Query(ctx, stmt, args)
	if err != nil {
		return nil, err
	}
	columns := make([]string, 0, len(rows.FieldDescriptions()))
	for _, fd := range rows.FieldDescriptions() {
		columns = append(columns, fd.Name)
	}

	// for each database row / record, a map with the column names and row values is added to the allMaps slice
	var allMaps []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}
		err = rows.Scan(pointers...)
		if err != nil {
			return nil, err
		}
		resultMap := make(map[string]interface{})
		for i, val := range values {
			//fmt.Printf("Adding key=%s val=%v\n", columns[i], val)
			resultMap[columns[i]] = val
		}
		allMaps = append(allMaps, resultMap)
	}
	return allMaps, nil
}

// InsertRows - вставка строк в целевую таблицу
func (r *Repository) InsertRows(ctx context.Context, rows []map[string]interface{}, dstSrcMap map[string]string, targetColumns []string, targetTable string) error {
	if len(rows) == 0 {
		return nil
	}

	sqlQueryTemplate := `
insert into %s(%s) values %s;`

	var args []interface{}
	rowsToInsert := strings.Builder{}
	// генерация запроса и списка аргументов
	for i, row := range rows {
		rowsToInsert.WriteString("(")
		var str string
		start := i * len(row)
		for j := start; j < start+len(row); j++ {
			str += fmt.Sprintf("$%d, ", j+1)
		}
		str = strings.TrimSuffix(str, ", ") + ")"
		if i != len(rows)-1 {
			str += ","
		}
		rowsToInsert.WriteString(str)

		// достаем из мапы с колонками нужную колонку с помощью мапы маршрутов
		for _, col := range targetColumns {
			args = append(args, row[dstSrcMap[col]])
		}
	}
	cols := strings.Join(targetColumns, ",")
	query := fmt.Sprintf(sqlQueryTemplate, targetTable, cols, rowsToInsert.String())

	_, err := r.dest.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) applyFilters(ctx context.Context, query sq.SelectBuilder, filters *domain.Filters) sq.SelectBuilder {
	if filters == nil {
		return query
	}

	if filters.Where != nil {
		query = query.Where(*filters.Where)
	}
	if filters.GroupBy != nil {
		query = query.GroupBy(*filters.GroupBy)
	}
	if filters.Having != nil {
		query = query.GroupBy(*filters.Having)
	}
	if filters.OrderBy != nil {
		query = query.OrderBy(*filters.OrderBy)
	}
	if filters.Limit != nil && *filters.Limit >= 0 {
		query.Limit(uint64(*filters.Limit))
	}
	if filters.Offset != nil && *filters.Offset >= 0 {
		query.Offset(uint64(*filters.Offset))
	}

	return query
}

//func (r *Repository) WithTransaction(ctx context.Context, fn func(tx Tx) error) (err error) {
//	return r.WithTx(ctx, pgx.TxOptions{}, fn)
//}
//
//func (r *Repository) WithTx(ctx context.Context, options pgx.TxOptions, fn func(tx Tx) error) (err error) {
//	tx, err := r.dest.BeginTx(ctx, options)
//	if err != nil {
//		return
//	}
//
//	defer func() {
//		if p := recover(); p != nil {
//			// a panic occurred, rollback and repanic
//			_ = tx.Rollback(ctx)
//			panic(p)
//		} else if err != nil {
//			// something went wrong, rollback
//			_ = tx.Rollback(ctx)
//		} else {
//			// all good, commit
//			err = tx.Commit(ctx)
//		}
//	}()
//
//	err = fn(tx)
//
//	return
//}

func psql() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
