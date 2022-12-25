package domain

import (
	"fmt"
	"strings"
)

type Table struct {
	Name       string
	Attributes []*Attribute `json:"attributes"`
	//PrimaryKey  []string
	//ForeignKeys map[string]string
}

func (t *Table) CreateTableSql() string {
	sqlString := strings.Builder{}

	fmt.Fprintf(&sqlString, "create table %s (", t.Name)
	for i, attr := range t.Attributes {
		fmt.Fprintf(&sqlString, "\"%s\" %s", attr.Name, attr.Type)

		if i != len(t.Attributes)-1 {
			sqlString.WriteByte(',')
		}
	}

	sqlString.WriteByte(')')
	sqlString.WriteByte(';')

	return sqlString.String()
}
