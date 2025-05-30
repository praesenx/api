package main

import (
	"github.com/gocanto/blog/boost"
	"github.com/gocanto/blog/env"
	"github.com/gocanto/blog/pkg"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
)

var environment *env.Environment
var validator *pkg.Validator

func init() {
	secrets, validate := boost.Spark("./.env")

	environment = secrets
	validator = validate
}

func main() {
	dbConnection := boost.MakeDbConnection(environment)
	logs := boost.MakeLogs(environment)
	adminUser := boost.MakeAdminUser(environment)
	localSentry := boost.MakeSentry(environment)

	defer (*logs).Close()
	defer (*dbConnection).Close()

	mux := http.NewServeMux()

	app := boost.MakeApp(mux, &boost.App{
		Validator:    validator,
		Logs:         logs,
		DbConnection: dbConnection,
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
