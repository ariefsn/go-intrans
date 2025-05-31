package db

import (
	"context"
	"database/sql"
	"embed"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

//go:embed migrations/*.sql
var sqlMigrations embed.FS

func InitDB(dsn string) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

func MigrateDB(db *bun.DB) (*migrate.Migrations, *migrate.Migrator, error) {
	migrations := migrate.NewMigrations()

	err := migrations.Discover(sqlMigrations)
	if err != nil {
		return nil, nil, err
	}

	migrator := migrate.NewMigrator(db, migrations)

	err = migrator.Init(context.Background())
	if err != nil {
		return nil, nil, err
	}

	_, err = migrator.Migrate(context.Background())
	if err != nil {
		return nil, nil, err
	}

	return migrations, migrator, nil
}
