package database

import "database/sql"

type Driver interface {
	Ping()
	Close() bool
	Driver() *sql.DB
}
