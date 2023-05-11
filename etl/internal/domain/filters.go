package domain

type Filters struct {
	Where   *string `json:"where,omitempty"`
	GroupBy *string `json:"group_by,omitempty"`
	Having  *string `json:"having,omitempty"`
	OrderBy *string `json:"order_by,omitempty"`
	Limit   *int    `json:"limit,omitempty"`
	Offset  *int    `json:"offset,omitempty"`
}
