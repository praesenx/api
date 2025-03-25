package main

import (
	"github.com/gocanto/blog/app/contracts"
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/support"
	"github.com/joho/godotenv"
	"strconv"
)

func getDatabaseConnection() *contracts.DatabaseDriver {
	dbConn, err := database.MakeConnection(environment)

	if err != nil {
		panic("DB: error connecting to PostgreSQL: " + err.Error())
	}

	return &dbConn
}

func getLogsDriver() *contracts.LogsDriver {
	lDriver, err := support.MakeDefaultFileLogs(environment.Logs)

	if err != nil {
		panic("Logs: error opening logs file: " + err.Error())
	}

	return &lDriver
}

func getEnvironment(validate support.Validator) support.Environment {
	errorSufix := "Environment: "

	values, err := godotenv.Read("./../.env")

	if err != nil {
		panic(errorSufix + "invalid .env file: " + err.Error())
	}

	port, _ := strconv.Atoi(values["ENV_DB_PORT"])
	portSecondary, _ := strconv.Atoi(values["ENV_DB_PORT_SECONDARY"])

	app := support.AppEnvironment{
		Name: values["ENV_APP_NAME"],
		Type: values["ENV_APP_ENV_TYPE"],
	}

	db := support.DBEnvironment{
		UserName:      values["ENV_DB_USER_NAME"],
		UserPassword:  values["ENV_DB_USER_PASSWORD"],
		DatabaseName:  values["ENV_DB_DATABASE_NAME"],
		Port:          port,
		PortSecondary: portSecondary,
		Host:          values["ENV_DB_HOST"],
		DriverName:    dbDriverName,
		BinDir:        values["EN_DB_BIN_DIR"],
		URL:           values["ENV_DB_URL"],
	}

	globalAdmin := support.GlobalAdmin{
		Salt:  values["ENV_APP_ADMIN_USER_TOKEN_SALT"],
		Token: values["ENV_APP_ADMIN_USER_TOKEN"],
	}

	logs := support.LogsEnvironment{
		Level:      values["ENV_APP_LOG_LEVEL"],
		Dir:        values["ENV_APP_LOGS_DIR"],
		DateFormat: values["ENV_APP_LOGS_DATE_FORMAT"],
	}

	net := support.NetEnvironment{
		HttpHost: values["ENV_HTTP_HOST"],
		HttpPort: values["ENV_HTTP_PORT"],
	}

	if _, err := validate.Rejects(app); err != nil {
		panic(errorSufix + "invalid app values: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(app); err != nil {
		panic(errorSufix + "invalid db values: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(app); err != nil {
		panic(errorSufix + "invalid global admin values: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(logs); err != nil {
		panic(errorSufix + "invalid logs values: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(net); err != nil {
		panic(errorSufix + "invalid network values: " + validate.GetErrorsAsJason())
	}

	env := support.Environment{
		App:     app,
		DB:      db,
		Admin:   globalAdmin,
		Logs:    logs,
		Network: net,
	}

	if _, err := validate.Rejects(env); err != nil {
		panic(errorSufix + "invalid env values: " + validate.GetErrorsAsJason())
	}

	return env
}
