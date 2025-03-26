package main

import (
	"github.com/go-playground/validator/v10"
	environment2 "github.com/gocanto/blog/app/environment"
	"github.com/gocanto/blog/app/support"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
)

const dbDriverName = "postgres"

var environment environment2.Environment
var verifier *support.Validator

func init() {
	verifier := support.MakeValidatorFrom(validator.New(
		validator.WithRequiredStructEnabled(),
	))

	environment = getEnvironment(*verifier)
}

func main() {
	dbConnection := getDatabaseConnection()
	logsDriver := getLogsDriver()

	defer (*logsDriver).Close()
	defer (*dbConnection).Close()

	mux := http.NewServeMux()
	router := getRouter(mux, logsDriver)

	router.registerUsers(dbConnection)

	slog.Info("GORM DSN :" + environment.DB.GetDSN())
	slog.Info("Starting new server on :" + environment.Network.HttpPort)

	if err := http.ListenAndServe(environment.Network.GetHostURL(), mux); err != nil {
		slog.Error("Error starting server", "error", err)
		panic("Error starting server.")
	}
}
