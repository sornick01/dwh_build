package domain

import "strings"

type Schema struct {
	Name   string  `json:"name"`
	Tables []Table `json:"tables"`
}

func (s *Schema) toSql(builder *strings.Builder) {
	builder.WriteString("create schema ")
	builder.WriteString(s.Name)
	builder.WriteString(";\n")

	for _, table := range s.Tables {
		table.toSql(builder, s.Name)
	}
}
