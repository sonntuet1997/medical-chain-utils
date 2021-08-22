package utils

type DBQuery interface {
	//UpdateOne update value where query true, if no row update, return e.ErrNoRowAffected
	UpdateOne(value interface{}, query *QueryStruct) error
	//InsertOne insert value to table, auto update ID with gonanoid.New()
	InsertOne(value interface{}) (interface{}, error)
	//Delete delete all values match query, if no row delete, return e.ErrNoRowAffected
	Delete(query *QueryStruct) error
	//SelectOne select the first value that match query, return it with fields, if no row return error
	SelectOne(query *QueryStruct, field ...interface{}) (interface{}, error)
	//SelectOne select all values that match query, return it with fields, if no row return error
	Select(query *QueryStruct, field ...interface{}) ([]interface{}, error)
	//Close close connection
	Close() error
	Migrate() error
}

type QueryStruct struct {
	Query interface{}
	Args  []interface{}
}

func NewQuery(query interface{}, args ...interface{}) *QueryStruct {
	return &QueryStruct{
		Query: query,
		Args:  args,
	}
}

//ParseField parse field to slice string
func ParseField(field []interface{}) []string {
	f := make([]string, 0)
	for _, m := range field {
		temp, _ := m.(string)
		f = append(f, temp)
	}
	return f
}
