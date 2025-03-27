package main

import (
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/support"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
)

const dbDriverName = "postgres"

var environment env.Environment
var validator *support.Validator

func init() {
	val := support.MakeValidatorFrom(baseValidator.New(
		baseValidator.WithRequiredStructEnabled(),
	))

	environment = resolveEnv(val)
	validator = val
}

func main() {
	orm := makeORM(&environment)
	logs := makeLogs(&environment)

	defer (*logs).Close()
	defer (*orm).Close()

	mux := http.NewServeMux()

	router := makeRouter(mux, &environment, &Container{
		logs:      logs,
		orm:       orm,
		validator: validator,
	})

	router.registerUsers()

	(*orm).Ping()
	slog.Info("Starting new server on :" + environment.Network.HttpPort)

	if err := http.ListenAndServe(environment.Network.GetHostURL(), mux); err != nil {
		slog.Error("Error starting server", "error", err)
		panic("Error starting server." + err.Error())
	}
}
