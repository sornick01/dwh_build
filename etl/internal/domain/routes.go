package domain

import (
	"fmt"
	"strings"
)

type Column struct {
	Name string  `json:"name"`
	As   *string `json:"as,omitempty"`
}

type Row struct {
	Schema  *string  `json:"schema,omitempty"`
	Table   string   `json:"table"`
	Columns []Column `json:"columns"`
}

type Route struct {
	Source      Row     `json:"source"`
	Destination Row     `json:"destination"`
	Condition   *string `json:"condition,omitempty"`
}

type Routes struct {
	Routes []Route `json:"routes,omitempty"`
}

func (r *Routes) ToSql() string {
	var builder strings.Builder
	r.buildRoutes(&builder)

	return builder.String()
}

func (r *Routes) buildRoutes(builder *strings.Builder) {
	if len(r.Routes) == 0 {
		return
	}

	builder.WriteString(`
set transaction isolation level repeatable read;
begin transaction;
`)
	for _, route := range r.Routes {
		route.buildRoute(builder)
	}
	builder.WriteString(`
commit;
`)
}

func (r *Route) buildRoute(builder *strings.Builder) {
	var dest string
	if r.Destination.Schema != nil {
		dest = fmt.Sprintf("%s.%s", *r.Destination.Schema, r.Destination.Table)
	} else {
		dest = fmt.Sprintf("%s", r.Destination.Table)
	}

	var src string
	if r.Destination.Schema != nil {
		src = fmt.Sprintf("%s.%s", *r.Source.Schema, r.Source.Table)
	} else {
		src = fmt.Sprintf("%s", r.Source.Table)
	}

	var destCols string
	for _, col := range r.Destination.Columns {
		destCols += fmt.Sprintf("%s, ", col)
	}
	destCols = strings.TrimSuffix(destCols, ", ")

	var srcCols string
	for _, col := range r.Source.Columns {
		if col.As != nil {
			srcCols += fmt.Sprintf("%s as %s, ", col.Name, *col.As)
		} else {
			srcCols += fmt.Sprintf("%s, ", col.Name)
		}
	}
	srcCols = strings.TrimSuffix(srcCols, ", ")

	str := fmt.Sprintf(`
insert into %s (%s)
select %s
from %s;
`, dest, destCols, srcCols, src)

	if r.Condition != nil {
		str = strings.TrimSuffix(str, `;
`)
		str += fmt.Sprintf(`
where %s;
`, *r.Condition)
	}
	builder.WriteString(str)
}
