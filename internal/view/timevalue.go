package view

type TimeValue struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

type Column struct {
	Type string
	Name string
}

type RowValue struct {
	Time  string
	Value float64
}
