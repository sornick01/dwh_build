package domain

import "strings"

type Table struct {
	Name       string      `json:"name"`
	Attributes []Attribute `json:"attributes"`
	Constraints
}

func (t *Table) ToSql(builder *strings.Builder, schemaName string) {
	builder.WriteString("create table ")
	builder.WriteString(schemaName)
	builder.WriteByte('.')
	builder.WriteString(t.Name)
	builder.WriteString(" (\n")
	for i, attribute := range t.Attributes {
		attribute.ToSql(builder)
		if i != len(t.Attributes)-1 {
			builder.WriteString(", ")
		}
		builder.WriteByte('\n')
	}
	builder.WriteString(");\n")
}
