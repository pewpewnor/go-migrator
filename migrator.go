package nomad

import (
	"database/sql"
	"fmt"
)

type Migration interface {
	Up(*Schema)
	Down(*Schema)
}

type Migrator struct {
	schema     *Schema
	migrations []Migration
}

func NewMigrator(driverName string, dataSourceName string) *Migrator {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("cannot connect to database: %v\n", err))
	}

	return &Migrator{schema: &Schema{db: db}}
}

func (migrator *Migrator) Close() {
	migrator.schema.db.Close()
	migrator.schema = nil
	migrator.migrations = nil
}

func (migrator *Migrator) AddMigration(migration ...Migration) {
	migrator.migrations = append(migrator.migrations, migration...)
}

func (migrator *Migrator) MigrateAllUp() {
	for _, migration := range migrator.migrations {
		migration.Up(migrator.schema)
	}
}

func (migrator *Migrator) MigrateAllDown() {
	for i := len(migrator.migrations) - 1; i >= 0; i-- {
		migrator.migrations[i].Down(migrator.schema)
	}
}

func (migrator *Migrator) MigrateUp(migrations ...Migration) {
	for _, migration := range migrations {
		migration.Up(migrator.schema)
	}
}

func (migrator *Migrator) MigrateDown(migrations ...Migration) {
	for _, migration := range migrations {
		migration.Down(migrator.schema)
	}
}
