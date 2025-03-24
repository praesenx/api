package main

import (
	"github.com/gocanto/blog/app/contracts"
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/support"
	"os"
	"strconv"
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

func mustInitialiseEnvironment() *Environment {
	port, _ := strconv.Atoi(os.Getenv("ENV_DB_PORT"))
	portSecondary, _ := strconv.Atoi(os.Getenv("ENV_DB_PORT_SECONDARY"))

	env := &Environment{
		appName:         os.Getenv("ENV_APP_NAME"),
		appEnv:          os.Getenv("ENV_APP_ENV"),
		appLogLevel:     os.Getenv("ENV_APP_LOG_LEVEL"),
		dbUserName:      os.Getenv("ENV_DB_USER_NAME"),
		dbUserPassword:  os.Getenv("ENV_DB_USER_PASSWORD"),
		dbDatabaseName:  os.Getenv("ENV_DB_DATABASE_NAME"),
		dbPort:          port,
		dbPortSecondary: portSecondary,
		dbHost:          os.Getenv("ENV_DB_HOST"),
		dbDriver:        os.Getenv("ENV_DB_DRIVER"),
		dbBinDir:        os.Getenv("EN_DB_BIN_DIR"), // Note the slight difference in the env var name
		dbURL:           os.Getenv("ENV_DB_URL"),
	}

	if env.hasErrors() {
		panic("error loading the .env file: ")
	}

	return env
}
