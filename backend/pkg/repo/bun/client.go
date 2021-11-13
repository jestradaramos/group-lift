package bun

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	DB *bun.DB
}

func InitBunDB(dsn string) *DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	return &DB{db}
}
