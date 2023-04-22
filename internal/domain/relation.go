package domain

import (
	"fmt"
	"strings"
)

type (
	RelationType string
	//OnActionType  string
	//OperationType string
)

const (
	OneToOne   RelationType = "one_to_one"
	OneToMany  RelationType = "one_to_many"
	ManyToMany RelationType = "many_to_many"
	//Cascade         OnActionType  = "cascade"
	//Restrict        OnActionType  = "restrict"
	//SetNull         OnActionType  = "set null"
	//SetDefault      OnActionType  = "set default"
	//NoAction        OnActionType  = "no action"
	//DeleteOperation OperationType = "delete"
	//UpdateOperation OperationType = "update"
)

type OnDeleteUpdate struct {
	Operation string   `json:"operation"`
	Action    string   `json:"action"`
	Columns   []string `json:"columns,omitempty"`
}

type RelationTable struct {
	Schema string `json:"schema"`
	Table  string `json:"table"`
	Field  string `json:"field"`
}

type Relation struct {
	ReferenceTable RelationTable    `json:"reference_table"` // кто ссылается
	ReferenceTo    RelationTable    `json:"reference_to"`    // куда ссылается
	RelationType   RelationType     `json:"relation_type"`
	OnDeleteUpdate []OnDeleteUpdate `json:"on_delete,omitempty"`
}

func (r *Relation) BuildRelationSql(builder *strings.Builder) {
	switch r.RelationType {
	case OneToOne:
		r.buildOneToOne(builder)
	case OneToMany:
		r.buildOneToMany(builder)
	case ManyToMany:
		r.buildManyToMany(builder)
	}
}

func (r *Relation) buildOneToOne(builder *strings.Builder) {
	var onAction string
	if r.OnDeleteUpdate != nil {
		onAction = fmt.Sprintf(` on %s %s`, r.OnDeleteUpdate[0].Operation, r.OnDeleteUpdate[0].Action)
	}
	str := fmt.Sprintf(`
alter table %s.%s
    add constraint fk_%s_%s foreign key (%s) references %s.%s (%s) %s;
`,
		r.ReferenceTable.Schema, r.ReferenceTable.Table, r.ReferenceTable.Table, r.ReferenceTo.Table, r.ReferenceTable.Field,
		r.ReferenceTo.Schema, r.ReferenceTo.Table, r.ReferenceTo.Field, onAction)

	builder.WriteString(str)
}

func (r *Relation) buildOneToMany(builder *strings.Builder) {
	var onAction string
	if r.OnDeleteUpdate != nil {
		onAction = fmt.Sprintf(` on %s %s`, r.OnDeleteUpdate[0].Operation, r.OnDeleteUpdate[0].Action)
	}
	str := fmt.Sprintf(`
alter table %s.%s
    add constraint fk_%s_%s foreign key (%s) references %s.%s (%s) %s;
`,
		r.ReferenceTable.Schema, r.ReferenceTable.Table, r.ReferenceTable.Table, r.ReferenceTo.Table, r.ReferenceTable.Field,
		r.ReferenceTo.Schema, r.ReferenceTo.Table, r.ReferenceTo.Field, onAction)

	builder.WriteString(str)
}

func (r *Relation) buildManyToMany(builder *strings.Builder) {
	var (
		onActionReferenceTable string
		onActionReferenceTo    string
	)
	if r.OnDeleteUpdate != nil {
		onActionReferenceTable = fmt.Sprintf(` on %s %s`, r.OnDeleteUpdate[0].Operation, r.OnDeleteUpdate[0].Action)
		onActionReferenceTo = fmt.Sprintf(` on %s %s`, r.OnDeleteUpdate[1].Operation, r.OnDeleteUpdate[1].Action)
	}

	str := fmt.Sprintf(`
create table %s.%s_%s
(
    "%s"    int references %s.%s (%s) %s,
    "%s"    int references %s.%s (%s) %s
);
`,
		r.ReferenceTable.Schema, r.ReferenceTable.Table, r.ReferenceTo.Table,
		r.ReferenceTable.Field, r.ReferenceTable.Schema, r.ReferenceTable.Table, r.ReferenceTable.Field, onActionReferenceTable,
		r.ReferenceTo.Field, r.ReferenceTo.Schema, r.ReferenceTo.Table, r.ReferenceTo.Field, onActionReferenceTo)

	builder.WriteString(str)
}
