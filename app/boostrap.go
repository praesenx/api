package main

import (
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/logger"
	"github.com/gocanto/blog/app/logger/filesmanager"
	"github.com/gocanto/blog/app/support"
	"github.com/joho/godotenv"
	"strconv"
)

func makeORM(env *env.Environment) *database.Driver {
	dbConn, err := database.MakeORM(env)

	if err != nil {
		panic("DB: error connecting to PostgreSQL: " + err.Error())
	}

	return &dbConn
}

func makeLogs(env *env.Environment) *logger.Managers {
	lDriver, err := filesmanager.MakeFilesManager(env)

	if err != nil {
		panic("Logs: error opening logs file: " + err.Error())
	}

	return &lDriver
}

func resolveEnv(validate *support.Validator) env.Environment {
	errorSufix := "Environment: "

	values, err := godotenv.Read("./../.env")

	if err != nil {
		panic(errorSufix + "invalid .env file: " + err.Error())
	}

	port, _ := strconv.Atoi(values["ENV_DB_PORT"])

	app := env.AppEnvironment{
		Name: values["ENV_APP_NAME"],
		Type: values["ENV_APP_ENV_TYPE"],
	}

	db := env.DBEnvironment{
		UserName:     values["ENV_DB_USER_NAME"],
		UserPassword: values["ENV_DB_USER_PASSWORD"],
		DatabaseName: values["ENV_DB_DATABASE_NAME"],
		Port:         port,
		Host:         values["ENV_DB_HOST"],
		DriverName:   dbDriverName,
		BinDir:       values["EN_DB_BIN_DIR"],
		URL:          values["ENV_DB_URL"],
		SSLMode:      values["ENV_DB_SSL_MODE"],
		TimeZone:     values["ENV_DB_TIMEZONE"],
	}

	globalAdmin := env.GlobalAdmin{
		Salt:  values["ENV_APP_ADMIN_USER_TOKEN_SALT"],
		Token: values["ENV_APP_ADMIN_USER_TOKEN"],
	}

	logs := env.LogsEnvironment{
		Level:      values["ENV_APP_LOG_LEVEL"],
		Dir:        values["ENV_APP_LOGS_DIR"],
		DateFormat: values["ENV_APP_LOGS_DATE_FORMAT"],
	}

	net := env.NetEnvironment{
		HttpHost: values["ENV_HTTP_HOST"],
		HttpPort: values["ENV_HTTP_PORT"],
	}

	//validate := verifier

	if _, err = validate.Rejects(app); err != nil {
		panic(errorSufix + "invalid app model: " + validate.GetErrorsAsJason())
	}

	if _, err = validate.Rejects(db); err != nil {
		panic(errorSufix + "invalid db model: " + validate.GetErrorsAsJason())
	}

	if _, err = validate.Rejects(globalAdmin); err != nil {
		panic(errorSufix + "invalid global admin model: " + validate.GetErrorsAsJason())
	}

	if _, err = validate.Rejects(logs); err != nil {
		panic(errorSufix + "invalid logs model: " + validate.GetErrorsAsJason())
	}

	if _, err = validate.Rejects(net); err != nil {
		panic(errorSufix + "invalid network model: " + validate.GetErrorsAsJason())
	}

	blog := env.Environment{
		App:     app,
		DB:      db,
		Admin:   globalAdmin,
		Logs:    logs,
		Network: net,
	}

	if _, err = validate.Rejects(blog); err != nil {
		panic(errorSufix + "invalid blog environment model: " + validate.GetErrorsAsJason())
	}

	return blog
}
