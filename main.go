package main

import (
	"github.com/getsentry/sentry-go"
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/gocanto/blog"
	"github.com/gocanto/blog/env"
	"github.com/gocanto/blog/webkit"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
)

var environment *env.Environment
var validator *webkit.Validator

func init() {
	val := webkit.MakeValidatorFrom(baseValidator.New(
		baseValidator.WithRequiredStructEnabled(),
	))

	values, err := godotenv.Read("./../.env")

	if err != nil {
		panic("invalid .env file: " + err.Error())
	}

	environment = blog.MakeEnv(values, val)
	validator = val
}

func main() {
	defer sentry.Recover()

	dbConnection := blog.MakeDbConnection(environment)
	logs := blog.MakeLogs(environment)
	adminUser := blog.MakeAdminUser(environment)
	localSentry := blog.MakeSentry(environment)

	defer (*logs).Close()
	defer (*dbConnection).Close()

	mux := http.NewServeMux()

	app := blog.MakeApp(mux, &blog.App{
		Validator:    validator,
		Logs:         logs,
		dbConnection: dbConnection,
		AdminUser:    adminUser,
		Env:          environment,
		Mux:          mux,
		Sentry:       localSentry,
	})

	app.RegisterUsers()

	(*dbConnection).Ping()
	slog.Info("Starting new server on :" + environment.Network.HttpPort)

	if err := http.ListenAndServe(environment.Network.GetHostURL(), mux); err != nil {
		slog.Error("Error starting server", "error", err)
		panic("Error starting server." + err.Error())
	}
}
