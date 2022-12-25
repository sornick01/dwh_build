package types_samples

type Database struct {
	Name   string
	Tables []Table
}

type Table struct {
	Name       string
	Attributes []Attribute
}

type Attribute struct {
	Name string
	Type string
}
