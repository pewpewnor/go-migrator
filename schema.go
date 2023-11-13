package nomad

import (
	"database/sql"
	"fmt"
	"log"
)

type Schema struct {
	db *sql.DB
}

func (schema *Schema) CreateTable(table *Table) {
	columns := ""

	for _, attr := range table.attributes {
		column := fmt.Sprintf("%v %v", attr.collumnName, attr.collumnType)
		if attr.primaryKey {
			column += " PRIMARY KEY"
		}
		if !attr.nullable {
			column += " NOT NULL"
		}

		if columns == "" {
			columns += column
		} else {
			columns += fmt.Sprintf(", %v", column)
		}
	}

	for _, constraint := range table.constraints {
		if columns == "" {
			columns += constraint
		} else {
			columns += fmt.Sprintf(", %v", constraint)
		}
	}

	log.Println(columns)

	_, err := schema.db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (%v)", table.tableName, columns))
	if err != nil {
		panic(fmt.Sprintf("Nomad cannot create table '%v': %v", table.tableName, err))
	}
}

func (schema *Schema) DropTable(tableName string) {
	_, err := schema.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %v", tableName))
	if err != nil {
		panic(fmt.Sprintf("Nomad cannot drop table '%v': %v", tableName, err))
	}
}

func (schema *Schema) DropTableCascade(tableName string) {
	_, err := schema.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %v CASCADE", tableName))
	if err != nil {
		panic(fmt.Sprintf("Nomad cannot drop table '%v': %v", tableName, err))
	}
}

func (schema *Schema) AlterTable(table *Table) {

}
