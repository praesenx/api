package main

import (
    "flag"
    "fmt"
    "github.com/gocanto/blog/bootstrap"
    "github.com/gocanto/blog/database"
    "github.com/gocanto/blog/database/seeder/seeds"
    "github.com/gocanto/blog/env"
    "github.com/gocanto/blog/webkit/cli"
    "os"
)

var environment *env.Environment
var textColour *cli.TextColour

func init() {
    secrets, _ := bootstrap.Spark("./.env")

    environment = secrets
}

func main() {
    dbConnection := bootstrap.MakeDbConnection(environment)
    logs := bootstrap.MakeLogs(environment)
    textColour = makeTextColour("Seeders are forbidden in production")

    defer (*logs).Close()
    defer (*dbConnection).Close()

    // ---- Path: 1
    if *wantsTruncation() == true {
        truncateDB(dbConnection, environment)
        textColour.SetMessage("DB successfully truncated.", cli.Green)
        return
    }

    // ---- Path: 2
    createUsers(dbConnection)

    fmt.Println("\nDone ...")
}

func makeTextColour(message string) *cli.TextColour {
    textColour, err := cli.MakePaddedTextColour(message, cli.Red)

    if err != nil {
        panic(err)
    }

    return textColour
}

func wantsTruncation() *bool {
    truncate := flag.Bool("truncate", false, "Truncate the database before seeding")
    flag.Parse()

    return truncate
}

func truncateDB(dbConnection *database.Connection, environment *env.Environment) {
    if environment.App.IsProduction() {
        fmt.Print(textColour.Get())
        os.Exit(1)
    }

    truncate := database.MakeTruncate(dbConnection, environment)

    if err := truncate.Execute(); err != nil {
        panic(err)
    }
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
