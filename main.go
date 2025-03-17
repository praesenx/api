package main

import (
	"fmt"
	"github.com/gocanto/blog/bootstrap"
	"github.com/gocanto/blog/users"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	_, err := bootstrap.LogInFile(
		fmt.Sprintf("./storage/logs/logs_%s.log", time.Now().UTC().Format("2006_02_01")),
	)

	if err != nil {
		slog.Error("Error initializing logs.", "error", err)
		panic("Error initializing logs.")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", users.Create)

	slog.Info("Starting server on :8080")
	err = http.ListenAndServe("localhost:8080", mux)

	if err != nil {
		slog.Error("Error starting server", "error", err)
		panic("Error starting server.")
	}
}
