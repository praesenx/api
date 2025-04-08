package blog

import (
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	blog2 "github.com/gocanto/blog"
	"github.com/gocanto/blog/database"
	env2 "github.com/gocanto/blog/env"
	"github.com/gocanto/blog/users"
	webkit2 "github.com/gocanto/blog/webkit"
	llogs2 "github.com/gocanto/blog/webkit/llogs"
	"log"
	"strconv"
	"strings"
	"time"
)

func MakeSentry(env *env2.Environment) *webkit2.Sentry {
	cOptions := sentry.ClientOptions{
		Dsn:   blog2.environment.Sentry.DSN,
		Debug: true,
	}

	if err := sentry.Init(cOptions); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)

	options := sentryhttp.Options{}
	handler := sentryhttp.New(options)

	return &webkit2.Sentry{
		Handler: handler,
		Options: &options,
		Env:     env,
	}
}

func MakeDbConnection(env *env2.Environment) *database.Connection {
	dbConn, err := database.MakeConnection(env)

	if err != nil {
		panic("Sql: error connecting to PostgreSQL: " + err.Error())
	}

	return dbConn
}

func MakeLogs(env *env2.Environment) *llogs2.Driver {
	lDriver, err := llogs2.MakeFilesLogs(env)

	if err != nil {
		panic("Logs: error opening logs file: " + err.Error())
	}

	return &lDriver
}

func MakeAdminUser(env *env2.Environment) *users.AdminUser {
	return &users.AdminUser{
		PublicToken:  env.App.AppUserAmin.PublicToken,
		PrivateToken: env.App.AppUserAmin.PrivateToken,
	}
}

func MakeEnv(values map[string]string, validate *webkit2.Validator) *env2.Environment {
	errorSufix := "Environment: "

	port, _ := strconv.Atoi(values["ENV_DB_PORT"])

	userAminEnvValues := &env2.AppUserAminEnvValues{
		PublicToken:  strings.Trim(values["ENV_APP_ADMIN_PUBLIC_TOKEN"], " "),
		PrivateToken: strings.Trim(values["ENV_APP_ADMIN_PRIVATE_TOKEN"], " "),
	}

	app := env2.AppEnvironment{
		Name:        values["ENV_APP_NAME"],
		Type:        values["ENV_APP_ENV_TYPE"],
		AppUserAmin: userAminEnvValues,
	}

	db := env2.DBEnvironment{
		UserName:     values["ENV_DB_USER_NAME"],
		UserPassword: values["ENV_DB_USER_PASSWORD"],
		DatabaseName: values["ENV_DB_DATABASE_NAME"],
		Port:         port,
		Host:         values["ENV_DB_HOST"],
		DriverName:   "postgres",
		BinDir:       values["EN_DB_BIN_DIR"],
		URL:          values["ENV_DB_URL"],
		SSLMode:      values["ENV_DB_SSL_MODE"],
		TimeZone:     values["ENV_DB_TIMEZONE"],
	}

	logsCreds := env2.LogsEnvironment{
		Level:      values["ENV_APP_LOG_LEVEL"],
		Dir:        values["ENV_APP_LOGS_DIR"],
		DateFormat: values["ENV_APP_LOGS_DATE_FORMAT"],
	}

	net := env2.NetEnvironment{
		HttpHost: values["ENV_HTTP_HOST"],
		HttpPort: values["ENV_HTTP_PORT"],
	}

	sentry := env2.SentryEnvironment{
		DSN: values["ENV_SENTRY_DSN"],
		CSP: values["ENV_SENTRY_CSP"],
	}

	if _, err := validate.Rejects(app); err != nil {
		panic(errorSufix + "invalid [APP] model: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(db); err != nil {
		panic(errorSufix + "invalid [Sql] model: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(userAminEnvValues); err != nil {
		panic(errorSufix + "invalid [AppUserAminEnvValues] model: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(logsCreds); err != nil {
		panic(errorSufix + "invalid [Logs Creds] model: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(net); err != nil {
		panic(errorSufix + "invalid [NETWORK] model: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(sentry); err != nil {
		panic(errorSufix + "invalid [SENTRY] model: " + validate.GetErrorsAsJason())
	}

	blog := &env2.Environment{
		App:     app,
		DB:      db,
		Logs:    logsCreds,
		Network: net,
		Sentry:  sentry,
	}

	if _, err := validate.Rejects(blog); err != nil {
		panic(errorSufix + "invalid blog [ENVIRONMENT] model: " + validate.GetErrorsAsJason())
	}

	return blog
}
