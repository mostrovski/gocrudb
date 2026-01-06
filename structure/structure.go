package structure

type FilterBy struct {
	Field    string
	Operator string
	Value    any
}

type SortBy struct {
	Field     string
	Direction string
}

type Paginate struct {
	Page    uint
	PerPage uint
}

type Conditions struct {
	Filters    []FilterBy
	Sorts      []SortBy
	Pagination Paginate
}
