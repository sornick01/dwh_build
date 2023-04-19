package domain

import (
	"fmt"
	"strings"
)

type Default struct {
	Column string      `json:"column"`
	Value  interface{} `json:"value"`
}

type Constraints struct {
	Schema  string    `json:"schema"`
	Table   string    `json:"table"`
	Unique  []string  `json:"unique,omitempty"`
	NotNull []string  `json:"notNull,omitempty"`
	Default []Default `json:"default,omitempty"`
}

func (c *Constraints) BuildConstraints(builder *strings.Builder) {
	if len(c.Unique) != 0 {
		c.buildUniqueConstraint(builder)
	}

	if len(c.NotNull) != 0 {
		c.buildNotNullConstraint(builder)
	}

	if len(c.Default) != 0 {
		c.buildDefaultConstraint(builder)
	}
}

func (c *Constraints) buildUniqueConstraint(builder *strings.Builder) {
	str := fmt.Sprintf(`
alter table %s.%s add constraint unique_constr unique (`, c.Schema, c.Table)
	for i := range c.Unique {
		str += fmt.Sprintf("%s,", c.Unique[i])
	}
	strings.TrimSuffix(str, ",")
	str += fmt.Sprintf(");\n")
	builder.WriteString(str)
}

func (c *Constraints) buildNotNullConstraint(builder *strings.Builder) {
	var str string
	for i := range c.NotNull {
		str += fmt.Sprintf(`
alter table %s.%s alter column %s set not null;
`, c.Schema, c.Table, c.NotNull[i])
	}
	builder.WriteString(str)
}

func (c *Constraints) buildDefaultConstraint(builder *strings.Builder) {
	var str string

	for _, def := range c.Default {
		switch v := def.Value.(type) {
		case int:
			str += fmt.Sprintf(`
alter table %s.%s alter column %s set default %d;
`, c.Schema, c.Table, def.Column, v)
		default:
			str += fmt.Sprintf(`
alter table %s.%s alter column %s set default '%s';
`, c.Schema, c.Table, def.Column, v)
		}
	}
	builder.WriteString(str)
}
