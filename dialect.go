package ray

import (
	"context"
)

const (
	MySQL    = "mysql"
	SQLite   = "sqlite3"
	Postgres = "postgres"
)

type ExecQuerier interface {
	Exec(ctx context.Context, query string, args, v interface{}) error

	Query(ctx context.Context, query string, args, v interface{}) error
}

type Driver interface {
	ExecQuerier

	Close() error

	Dialect() string
}
