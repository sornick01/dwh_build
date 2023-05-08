package domain

import (
	"fmt"
	"strings"
)

type Table struct {
	Name       string      `json:"name"`
	Attributes []Attribute `json:"attributes"`
	PrimaryKey []string    `json:"primary_key,omitempty"`
}

func (t *Table) toSql(builder *strings.Builder, schemaName string) {
	dropIfExists := fmt.Sprintf(`drop table if exists %s.%s`, schemaName, t.Name)
	builder.WriteString(dropIfExists)

	builder.WriteString("create table ")
	builder.WriteString(schemaName)
	builder.WriteByte('.')
	builder.WriteString(t.Name)
	builder.WriteString("\n(\n")

	for i, attribute := range t.Attributes {
		attribute.toSql(builder)
		if i != len(t.Attributes)-1 {
			builder.WriteString(",\n")
		}
	}

	if len(t.PrimaryKey) != 0 {
		builder.WriteString(",\n\tprimary key (")
		for i, column := range t.PrimaryKey {
			builder.WriteString(column)
			if i != len(t.PrimaryKey)-1 {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(")\n")
	} else {
		builder.WriteByte('\n')
	}

	builder.WriteString(");\n")
}
