package nomad

type attribute struct {
	collumnName string
	collumnType string
	primaryKey  bool
	nullable    bool
}

func NewAttribute(columnName string, columnType string) *attribute {
	return &attribute{
		collumnName: columnName,
		collumnType: columnType,
		primaryKey:  false,
		nullable:    false,
	}
}
