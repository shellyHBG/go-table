package table

type table struct {
	name    string // table name
	content map[string]map[string]string
}

func (t *table) Get(key, subKey string) string {
	return t.content[key][subKey]
}

func (t *table) GetAll(key string) map[string]string {
	return t.content[key]
}
