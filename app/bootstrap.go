package main

import (
	"github.com/gocanto/blog/app/contracts"
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/support"
	"os"
)

func mustInitialiseDatabase() *contracts.DatabaseDriver {
	dbConn, err := database.MakeConnection("postgres", os.Getenv("ENV_DB_URL"))

	if err != nil {
		panic("error connecting to PostgreSQL: " + err.Error())
	}

	return &dbConn
}

func mustInitialiseLogsDriver() *contracts.LogsDriver {
	lDriver, err := support.MakeDefaultFileLogs()

	if err != nil {
		panic("error opening logs file: " + err.Error())
	}

	return &lDriver
}

func mustInitialiseValidator() *support.Validator {
	return support.MakeValidator()
}
