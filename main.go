package main

import (
    "github.com/gocanto/blog/bootstrap"
    "github.com/gocanto/blog/env"
    "github.com/gocanto/blog/pkgs"
    _ "github.com/lib/pq"
    "log/slog"
    "net/http"
)

var environment *env.Environment
var validator *pkgs.Validator

func init() {
    secrets, validate := bootstrap.Spark("./.env")

    environment = secrets
    validator = validate
}

func main() {
    dbConnection := bootstrap.MakeDbConnection(environment)
    logs := bootstrap.MakeLogs(environment)
    adminUser := bootstrap.MakeAdminUser(environment)
    localSentry := bootstrap.MakeSentry(environment)

    defer (*logs).Close()
    defer (*dbConnection).Close()

    mux := http.NewServeMux()

    app := bootstrap.MakeApp(mux, &bootstrap.App{
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
