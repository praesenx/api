package database

import (
	"database/sql"
	"fmt"
	"github.com/gocanto/blog/app/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

type Connection struct {
	url        string
	driverName string
	driver     *gorm.DB
	env        env.Environment
}

func MakeDbConnection(env env.Environment) (Driver, error) {
	dbEnv := env.DB
	driver, err := gorm.Open(postgres.Open(dbEnv.GetDSN()), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Connection{
		url:        dbEnv.URL,
		driver:     driver,
		driverName: dbEnv.DriverName,
		env:        env,
	}, nil
}

func (receiver *Connection) Close() bool {
	if sqlDB, err := receiver.driver.DB(); err != nil {
		slog.Error("There was an error closing the db: " + err.Error())

		return false
	} else {
		_ = sqlDB.Close()
	}

	return true
}

func (receiver *Connection) Ping() {
	var driver *sql.DB

	fmt.Println("\n---------")

	if conn, err := receiver.driver.DB(); err != nil {
		fmt.Println(fmt.Sprintf("error retrieving the db driver: %v", err.Error()))

		return
	} else {
		driver = conn
		fmt.Println(fmt.Sprintf("db driver adquired: %T", driver))
	}

	if err := driver.Ping(); err != nil {
		slog.Error("error pinging the db driver: " + err.Error())
	}

	fmt.Println(fmt.Sprintf("db driver is healthy: %+v", driver.Stats()))

	fmt.Println("---------")
}
func (receiver *Connection) Driver() *gorm.DB {
	return receiver.driver
}
