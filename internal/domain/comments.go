package domain

import (
	"fmt"
	"strings"
)

const (
	databaseType   = "database"
	schemaType     = "schema"
	tableType      = "table"
	columnType     = "column"
	constraintType = "constraint"
	//indexType      = "index"
)

type Comment struct {
	EntityName       string  `json:"entity_name"`
	EntityType       string  `json:"entity_type"`
	EntityExtraField *string `json:"entity_extra_field,omitempty"`
	Text             string  `json:"text"`
}

func (c *Comment) BuildComment(builder *strings.Builder) {
	var str string
	switch c.EntityType {
	case databaseType, schemaType, tableType, columnType:
		str = fmt.Sprintf(`comment on %s `, c.EntityType)
	case constraintType:
		if c.EntityExtraField != nil {
			str = fmt.Sprintf(`comment on %s on %s `, c.EntityType, c.EntityExtraField)
		}
	}

	str += fmt.Sprintf(`%s;
`, c.EntityName)
	builder.WriteString(str)
}
