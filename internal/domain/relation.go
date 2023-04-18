package domain

type RelationType int

const (
	OneToMany = iota
	OneToOne
	ManyToMany
)

type Relation struct {
	Referenced   string       `json:"referenced"`
	Reference    string       `json:"reference"`
	RelationType RelationType `json:"relation_type"`
}

//func NewRelation() Relation {
//
//}
