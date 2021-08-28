package utils

//ParseField parse field to slice string
func ParseField(field []interface{}) []string {
	f := make([]string, 0)
	for _, m := range field {
		temp, _ := m.(string)
		f = append(f, temp)
	}
	return f
}
