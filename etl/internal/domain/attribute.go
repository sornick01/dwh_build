package domain

import "strings"

type Attribute struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (a *Attribute) toSql(builder *strings.Builder) {
	builder.WriteByte('\t')
	builder.WriteString(a.Name)
	builder.WriteByte(' ')
	builder.WriteString(a.Type)
}
