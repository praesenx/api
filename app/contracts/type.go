package contracts

import "database/sql"

type LogsDriver interface {
	Close() bool
}

type DatabaseDriver interface {
	Ping()
	Close() bool
	Driver() *sql.DB
}
