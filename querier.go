package utils

type Querier interface {
	ToQuery() string
}
