package structure

type Condition struct {
	Field    string
	Operator string
	Value    any
}

type SortBy struct {
	Field     string
	Direction string
}
