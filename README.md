# Go Migrator

A small Go library for defining and applying database schema migrations. Migrations implement an Up and Down operation, while the schema helper provides basic table creation and removal with UUID and string columns, nullability, primary keys, and foreign-key constraints.

This project is an early implementation: it does not maintain a migration history table, and schema alteration support is not implemented. Register migrations in the order they should be applied; rollback runs them in reverse order.

## Requirements

- Go 1.21 or newer
- A database/sql driver for the database you want to use. The example below uses PostgreSQL and lib/pq.

## Example

~~~go
package main

import (
    "github.com/pewpewnor/nomad"
    _ "github.com/lib/pq"
)

type CreateUsers struct{}

func (CreateUsers) Up(schema *nomad.Schema) {
    table := schema.NewTable("users")
    table.UUID("id").PrimaryKey()
    table.String("name")
    schema.CreateTable(table)
}

func (CreateUsers) Down(schema *nomad.Schema) {
    schema.DropTable("users")
}

func main() {
    migrator := nomad.NewMigrator("postgres", "postgres://user@localhost:5432/app?sslmode=disable")
    defer migrator.Close()

    migrator.AddMigration(CreateUsers{})
    migrator.MigrateAllUp()
}
~~~

The module currently declares the import path `github.com/pewpewnor/nomad`. Supply a valid database connection string for your driver. The API is intentionally small and may change while the library is under development.
