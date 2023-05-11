package domain

import (
	"fmt"
)

//type Column struct {
//	NameOrValue string  `json:"name_or_value"`
//	As          *string `json:"as,omitempty"`
//}

//type Row struct {
//	//Schema *string `json:"schema,omitempty"`
//	Table  string  `json:"table"`
//	Column string  `json:"column"`
//	As     *string `json:"as,omitempty"`
//}

//type Additions struct {
//	Where   *string `json:"where,omitempty"`
//	GroupBy *string `json:"groupBy,omitempty"`
//	Having  *string `json:"having,omitempty"`
//	OrderBy *string `json:"orderBy,omitempty"`
//	Limit   *string `json:"limit,omitempty"`
//	Offset  *string `json:"offset,omitempty"`
//}

// Route - маршрут одной колонки
type Route struct {
	Source      string `json:"source"`      // Колонка источник
	Destination string `json:"destination"` // Колонка назначение
	//Additions   *Additions `json:"additions,omitempty"`
}

// TargetTable - целевая таблица
type TargetTable struct {
	Schema *string `json:"schema,omitempty"`
	Name   string  `json:"name"`
}

// Routes - маршруты
type Routes struct {
	//DatabaseName string      `json:"database_name"`
	SourceSQL   string      `json:"source_sql"`        // Источник данных в виде sql выражения (может быть подзапрос)
	TargetTable TargetTable `json:"target_table"`      // Целевая таблица
	Filters     *Filters    `json:"filters,omitempty"` // Фильтры
	Routes      []Route     `json:"routes"`            // Маршруты колонок
}

// String - вывод наименования целевой таблицы
func (t *TargetTable) String() string {
	if t.Schema != nil {
		return fmt.Sprintf("%s.%s", *t.Schema, t.Name)
	}
	return t.Name
}

// ListSrcColumns - получение списка столбцов-источников
func (r *Routes) ListSrcColumns() []string {
	list := make([]string, 0, len(r.Routes))

	for _, route := range r.Routes {
		list = append(list, route.Source)
	}
	return list
}

// ListTargetColumns - получение списка столбцов-назначений
func (r *Routes) ListTargetColumns() []string {
	list := make([]string, 0, len(r.Routes))

	for _, route := range r.Routes {
		list = append(list, route.Destination)
	}
	return list
}

// GetRouteMap - получение карты маршрутов столбцов
func (r *Routes) GetRouteMap() map[string]string {
	res := make(map[string]string)

	for _, route := range r.Routes {
		res[route.Destination] = route.Source
	}
	return res
}

//func (r *Routes) ToSql() string {
//	var builder strings.Builder
//	r.buildRoutes(&builder)
//
//	return builder.String()
//}

//func (r *Routes) buildRoutes(builder *strings.Builder) {
//	if len(r.Routes) == 0 {
//		return
//	}
//
//	builder.WriteString(`
//set transaction isolation level repeatable read;
//begin transaction;
//`)
//	for _, route := range r.Routes {
//		route.buildRoute(builder)
//	}
//	builder.WriteString(`
//commit;
//`)
//}

//func (r *Route) buildRoute(builder *strings.Builder) {
//	var dest string
//	if r.Destination.Schema != nil {
//		dest = fmt.Sprintf("%s.%s", *r.Destination.Schema, r.Destination.Table)
//	} else {
//		dest = fmt.Sprintf("%s", r.Destination.Table)
//	}
//
//	var src string
//	if r.Destination.Schema != nil {
//		src = fmt.Sprintf("%s.%s", *r.Source.Schema, r.Source.Table)
//	} else {
//		src = fmt.Sprintf("%s", r.Source.Table)
//	}
//
//	var destCols string
//	for _, col := range r.Destination.Columns {
//		destCols += fmt.Sprintf("%s, ", col)
//	}
//	destCols = strings.TrimSuffix(destCols, ", ")
//
//	var srcCols string
//	for _, col := range r.Source.Columns {
//		if col.As != nil {
//			srcCols += fmt.Sprintf("%s as %s, ", col.NameOrValue, *col.As)
//		} else {
//			srcCols += fmt.Sprintf("%s, ", col.NameOrValue)
//		}
//	}
//	srcCols = strings.TrimSuffix(srcCols, ", ")
//
//	// TODO: переделать на select
//	str := fmt.Sprintf(`
//insert into %s (%s)
//select %s
//from %s;
//`, dest, destCols, srcCols, src)
//
//	if r.Additions != nil {
//		str = strings.TrimSuffix(str, `;
//`)
//		str += r.Additions.build()
//	}
//
//	builder.WriteString(str)
//}

//func (a *Additions) build() string {
//	var str string
//
//	if a.Where != nil {
//		str += fmt.Sprintf("where %s", *a.Where)
//	}
//	if a.GroupBy != nil {
//		str += fmt.Sprintf("group by %s", *a.GroupBy)
//	}
//	if a.Having != nil {
//		str += fmt.Sprintf("having %s", *a.Having)
//	}
//	if a.OrderBy != nil {
//		str += fmt.Sprintf("order by %s", *a.OrderBy)
//	}
//	if a.Limit != nil {
//		str += fmt.Sprintf("limit %s", *a.Limit)
//	}
//	if a.Offset != nil {
//		str += fmt.Sprintf("offset %s", *a.Offset)
//	}
//
//	return fmt.Sprintf(`%s;
//`, str)
//}
