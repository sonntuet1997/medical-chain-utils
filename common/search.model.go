package common

type DefaultSearchModel struct {
	Skip      int
	Limit     int
	OrderBy   string
	OrderType string
	Fields    []string
}
