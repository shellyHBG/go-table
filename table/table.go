package table

type Table struct {
	name    string // table name
	content map[string]map[string]string
}

func (t *Table) GetName() string {
	return t.name
}

func (t *Table) Get(key, subKey string) string {
	return t.content[key][subKey]
}

func (t *Table) GetAll(key string) map[string]string {
	return t.content[key]
}
