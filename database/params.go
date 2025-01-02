package database

type sortKey int

const (
	ASC sortKey = iota
	DESC
)

type Params struct {
	key   string
	value interface{}
}

type SortParams struct {
	key   string
	value sortKey
}

var sortKeyString = map[sortKey]string{
	ASC:  "asc",
	DESC: "desc",
}

func (s sortKey) String() string {
	return sortKeyString[s]
}

func SetSort(key string, value sortKey) SortParams {
	return SortParams{key: key, value: value}
}

func SetParam(key string, value interface{}) Params {
	return Params{key: key, value: value}
}

func (s Params) GetKey() string {
	return s.key
}
func (s Params) GetValue() interface{} {
	return s.value
}

func (s SortParams) GetKey() string {
	return s.key
}
func (s SortParams) GetValue() sortKey {
	return s.value
}
