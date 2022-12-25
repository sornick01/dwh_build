package domain

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Database struct {
	Name string `json:"name"`
	//Tables map[string]*Table `json:"tables"`
	Tables []*Table `json:"tables"`
	Pool   *pgxpool.Pool
}

func (d *Database) CreateDatabase() string { //TODO: сразу вызов pool.Exec
	var sqlString string

	for _, table := range d.Tables {
		sqlString += table.CreateTableSql()
	}

	_, err := d.Pool.Exec(context.Background(), sqlString)
	if err != nil {
		log.Fatal(err)
	}

	return sqlString
}
