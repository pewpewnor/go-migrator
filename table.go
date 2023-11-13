package nomad

import "fmt"

type Table struct {
	tableName   string
	attributes  []*attribute
	constraints []string
}

func (table *Table) addAttribute(columnName string, columnType string) *Column {
	attr := NewAttribute(columnName, columnType)
	table.attributes = append(table.attributes, attr)

	return &Column{attr: attr}
}

func (table *Table) addExtraConstraint(constraint string) {
	table.constraints = append(table.constraints, constraint)
}

func (table *Table) UUID(columnName string) *Column {
	return table.addAttribute(columnName, "UUID")
}

func (table *Table) Varchar(columnName string, length uint16) *Column {
	return table.addAttribute(columnName, fmt.Sprintf("VARCHAR(%v)", length))
}

func (table *Table) String(columnName string) *Column {
	return table.Varchar(columnName, 255)
}

func (table *Table) ForeignKey(columnName string, foreignTableName string, foreignColumnName string) {
	table.addExtraConstraint(
		fmt.Sprintf("FOREIGN KEY (%v) REFERENCES %v(%v)",
			columnName, foreignTableName, foreignColumnName))
}
