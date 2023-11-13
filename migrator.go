package nomad

import (
	"database/sql"
	"fmt"
	"log"
)

type Migration interface {
	Up(*Migrator)
	Down(*Migrator)
}

type Migrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) *Migrator {
	return &Migrator{db: db}
}

func (migrator *Migrator) CreateTable(table *Table) {
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

	_, err := migrator.db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (%v)", table.tableName, columns))
	if err != nil {
		panic(fmt.Sprintf("Nomad cannot create table '%v': %v", table.tableName, err))
	}
}

func (migrator *Migrator) DropTable(tableName string) {
	_, err := migrator.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %v", tableName))
	if err != nil {
		panic(fmt.Sprintf("Nomad cannot drop table '%v': %v", tableName, err))
	}
}

func (migrator *Migrator) DropTableCascade(tableName string) {
	_, err := migrator.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %v CASCADE", tableName))
	if err != nil {
		panic(fmt.Sprintf("Nomad cannot drop table '%v': %v", tableName, err))
	}
}

func (migrator *Migrator) AlterTable(table *Table) {

}

func MigrateUp(migrator *Migrator, migrations ...Migration) {
	for _, migration := range migrations {
		migration.Up(migrator)
	}
}

func MigrateDown(migrator *Migrator, migrations ...Migration) {
	for _, migration := range migrations {
		migration.Down(migrator)
	}
}
