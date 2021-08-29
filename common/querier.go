package common

type Querier interface {
	ToQuery() string
}
