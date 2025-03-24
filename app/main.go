package main

import (
	"github.com/gocanto/blog/app/support"
	"github.com/gocanto/blog/app/user"
	"log/slog"
	"net/http"
)

func main() {
	if fileLogs, err := support.MakeDefaultFileLogs(); err != nil {
		panic("error opening file: " + err.Error())
	} else {
		defer fileLogs.Close()
	}

	validator := support.MakeValidator()

	users := user.Controller{
		Validator: validator,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", users.Create)

	slog.Info("Starting new server on :8080")

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		slog.Error("Error starting server", "error", err)
		panic("Error starting server.")
	}
}
