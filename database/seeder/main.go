package main

import (
    "fmt"
    "github.com/gocanto/blog/bootstrap"
    "github.com/gocanto/blog/database"
    "github.com/gocanto/blog/database/seeder/seeds"
    "github.com/gocanto/blog/env"
    "github.com/gocanto/blog/webkit/cli"
    "os"
)

var environment *env.Environment

func init() {
    secrets, _ := bootstrap.Spark("./.env")

    environment = secrets
}

func main() {
    dbConnection := bootstrap.MakeDbConnection(environment)
    logs := bootstrap.MakeLogs(environment)

    defer (*logs).Close()
    defer (*dbConnection).Close()

    // -- Truncate
    truncateDB(environment)
    // --

    dbConnection.Sql().Exec("SET session_replication_role = 'replica';")
    createUsers(dbConnection)
    dbConnection.Sql().Exec("SET session_replication_role = 'origin';")

    fmt.Println("\nDone ...")
}

func truncateDB(env *env.Environment) {
    text, err := cli.MakePaddedTextColour("Senders are forbidden in production.", cli.Red)

    if err != nil {
        panic(err)
    }

    if env.App.IsProduction() {
        fmt.Print(text.Get())
        os.Exit(1)
    }

    fmt.Println("OK.")
}

func createUsers(DBconn *database.Connection) {
    seeds.CreateUser(seeds.CreateUsersAttrs{
        DB:       DBconn,
        Username: "gocanto",
        Name:     "Gus",
        IsAdmin:  true,
    })

    seeds.CreateUser(seeds.CreateUsersAttrs{
        DB:       DBconn,
        Username: "liane",
        Name:     "Li",
        IsAdmin:  false,
    })
}
