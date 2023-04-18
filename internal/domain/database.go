package domain

import "strings"

type Database struct {
	Name      string     `json:"name"`
	Schemas   []Schema   `json:"schemas"`
	Relations []Relation `json:"relations"`
}

func (db *Database) ToSql() string {
	var builder strings.Builder

	for _, schema := range db.Schemas {
		schema.toSql(&builder)
	}

	for _, relation := range db.Relations {
		relation.BuildRelationSql(&builder)
	}

	return builder.String()
}
