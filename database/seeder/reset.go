package main

import (
    "github.com/gocanto/blog/database"
    "github.com/gocanto/blog/env"
)

func truncateDB(dbConnection *database.Connection, environment *env.Environment) {
    if environment.App.IsProduction() {
        panic(textColour.Print())
    }

    truncate := database.MakeTruncate(dbConnection, environment)

    if err := truncate.Execute(); err != nil {
        panic(err)
    }
}
