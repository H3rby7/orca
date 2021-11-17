package main

import (
	"database/sql"
	"orca/api"
	"orca/env"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	dsn := "postgres://orca:still-cool@localhost:5432/orca?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())

	env := &env.Env{
		DB: db,
	}
	api.SetUpRoutes(env)
}
