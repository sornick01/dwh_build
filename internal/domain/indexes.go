package domain

import (
	"fmt"
	"strings"
)

type Index struct {
	Name       string  `json:"name"`
	Unique     bool    `json:"unique,omitempty"`
	IfNotExist bool    `json:"ifNotExist,omitempty"`
	Table      string  `json:"table"`
	Column     string  `json:"column"`
	Using      *string `json:"using,omitempty"`
	With       *string `json:"with,omitempty"`
}

func (i *Index) BuildIndex(builder *strings.Builder) {
	var (
		unique      string
		ifNotExists string
		using       string
		with        string
	)
	if i.Unique {
		unique = "unique "
	}
	if i.IfNotExist {
		ifNotExists = "if not exists "
	}
	if i.Using != nil {
		using = *i.Using
	}
	if i.With != nil {
		with = *i.With
	}

	str := fmt.Sprintf(`create %sindex %s%s on %s %s(%s)%s;
`, unique, ifNotExists, i.Name, i.Table, using, i.Column, with)

	builder.WriteString(str)
}
