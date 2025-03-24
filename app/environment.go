package main

type Environment struct {
	appName         string
	appEnv          string
	appLogLevel     string
	dbUserName      string
	dbUserPassword  string
	dbDatabaseName  string
	dbPort          int
	dbPortSecondary int
	dbHost          string
	dbDriver        string
	dbBinDir        string // Note the slight difference in the env var name
	dbURL           string
}

func (receiver Environment) hasErrors() bool {
	return false
}
