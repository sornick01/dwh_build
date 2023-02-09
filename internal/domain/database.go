package domain

import "strings"

type Database struct {
	Name    string   `json:"name"`
	Schemas []Schema `json:"schemas"`
}

func (db *Database) ToSql() string {
	var builder strings.Builder

	for _, schema := range db.Schemas {
		schema.toSql(&builder)
	}

	return builder.String()
}
