package main

import (
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/support"
	"github.com/gocanto/blog/app/user"
	"log/slog"
	"net/http"
)

var validator *support.Validator
var dbConn *database.Connection

func init() {
	validator = support.MakeValidator()
	dbConn = &database.Connection{}
}

func main() {
	if fileLogs, err := support.MakeDefaultFileLogs(); err != nil {
		panic("error opening file: " + err.Error())
	} else {
		defer fileLogs.Close()
	}

	mux := http.NewServeMux()

	registerUsersFor(mux)

	slog.Info("Starting new server on :8080")

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		slog.Error("Error starting server", "error", err)
		panic("Error starting server.")
	}
}

func registerUsersFor(mux *http.ServeMux) {
	repository := user.NewRepository(*dbConn)

	service := user.NewProvider(repository, validator)
	_ = service.Register(mux)
}
