package main

import (
    "github.com/getsentry/sentry-go"
    "github.com/gocanto/blog/boostrap"
    "github.com/gocanto/blog/env"
    "github.com/gocanto/blog/webkit"
    _ "github.com/lib/pq"
    "log/slog"
    "net/http"
)

var environment *env.Environment
var validator *webkit.Validator

func init() {
    secrets, validate := boostrap.Spark("./.env")

    environment = secrets
    validator = validate
}

func main() {
    defer sentry.Recover()

    dbConnection := boostrap.MakeDbConnection(environment)
    logs := boostrap.MakeLogs(environment)
    adminUser := boostrap.MakeAdminUser(environment)
    localSentry := boostrap.MakeSentry(environment)

    defer (*logs).Close()
    defer (*dbConnection).Close()

    mux := http.NewServeMux()

    app := boostrap.MakeApp(mux, &boostrap.App{
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
