package domain

import "strings"

type Database struct {
	Name        string        `json:"name"`
	Schemas     []Schema      `json:"schemas"`
	Relations   []Relation    `json:"relations,omitempty"`
	Constraints []Constraints `json:"constraints,omitempty"`
	Indexes     []Index       `json:"indexes,omitempty"`
	Comments    []Comment     `json:"comments,omitempty"`
}

func (db *Database) ToSql() string {
	var builder strings.Builder

	for _, schema := range db.Schemas {
		schema.toSql(&builder)
	}

	for _, relation := range db.Relations {
		relation.BuildRelationSql(&builder)
	}

	for _, constraint := range db.Constraints {
		constraint.BuildConstraints(&builder)
	}

	for _, index := range db.Indexes {
		index.BuildIndex(&builder)
	}

	for _, comment := range db.Comments {
		comment.BuildComment(&builder)
	}

	return builder.String()
}
