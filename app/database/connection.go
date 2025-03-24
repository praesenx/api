package database

import (
	"database/sql"
	"fmt"
	"github.com/gocanto/blog/app/contracts"
	"log/slog"
)

type Connection struct {
	url    string
	driver *sql.DB
}

func MakeConnection(driverName string, url string) (contracts.DatabaseDriver, error) {
	driver, err := sql.Open(driverName, url)

	if err != nil {
		return nil, err
	}

	return &Connection{
		url:    url,
		driver: driver,
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
