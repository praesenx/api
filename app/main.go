package main

import (
	"github.com/gocanto/blog/app/contracts"
	"github.com/gocanto/blog/app/support"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
)

var validator *support.Validator
var dbConnection *contracts.DatabaseDriver
var logsDriver *contracts.LogsDriver

func init() {
	validator = mustInitialiseValidator()
	dbConnection = mustInitialiseDatabase()
	logsDriver = mustInitialiseLogsDriver()
}

func main() {
	defer (*logsDriver).Close()
	defer (*dbConnection).Close()

	mux := http.NewServeMux()

	usersRoutes(mux)

	slog.Info("Starting new server on :8080")

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		slog.Error("Error starting server", "error", err)
		panic("Error starting server.")
	}
}
