package database

import (
	"database/sql"
	"fmt"
	"github.com/gocanto/blog/app/contracts"
	"github.com/gocanto/blog/app/environment"
	"log/slog"
)

type Connection struct {
	url         string
	driver      *sql.DB
	driverName  string
	environment environment.Environment
}

func MakeConnection(environment environment.Environment) (contracts.DatabaseDriver, error) {
	dbEnv := environment.DB
	driver, err := sql.Open(dbEnv.DriverName, dbEnv.URL)

	if err != nil {
		return nil, err
	}

	return &Connection{
		url:         dbEnv.URL,
		driver:      driver,
		driverName:  dbEnv.DriverName,
		environment: environment,
	}, nil
}

func (receiver *Connection) Close() bool {
	if err := receiver.driver.Close(); err != nil {
		slog.Error("There was an error closing the db: " + err.Error())

		return false
	}

	return true
}

func (receiver *Connection) Ping() {
	if err := receiver.driver.Ping(); err != nil {
		fmt.Println(fmt.Sprintf("There was an error pinging the db: %v", err.Error()))
	}

	fmt.Println("DB Connected ....")
}
func (receiver *Connection) Driver() *sql.DB {
	return receiver.driver
}
