package types_samples

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

type (
	Database struct {
		Name   string            `json:"name"`
		Tables map[string]*Table `json:"tables"`
		Pool   *pgxpool.Pool
	}

	Table struct {
		//Name        string
		Attributes  map[string]*Attribute `json:"attributes"`
		PrimaryKey  []string
		ForeignKeys map[string]string
	}

	Attribute struct {
		//Name string
		Type string `json:"type"`
	}

	Relation struct {
	}
)

func (d *Database) CreateDatabase() string {
	var sqlString string
	for name, elem := range d.Tables {
		sqlString += fmt.Sprintf(elem.CreateTable(d.Pool), name) + ";\n"
	}

	return sqlString
}

func (t *Table) CreateTable(pool *pgxpool.Pool) string {
	sqlString := `CREATE TABLE %s (`
	for i, elem := range t.Attributes {
		sqlString = sqlString + i + elem.ToSql() + ","
	}
	sqlString = strings.TrimSuffix(sqlString, ",")
	//pool.Exec(context.Background(), "")
	return sqlString + ")"
}

func (a *Attribute) ToSql() string {
	return fmt.Sprintf(" %s", a.Type)
}
