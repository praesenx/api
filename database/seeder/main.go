package main

import (
    "fmt"
    "github.com/gocanto/blog/bootstrap"
    "github.com/gocanto/blog/database"
    "github.com/gocanto/blog/database/seeder/seed"
    "github.com/gocanto/blog/env"
    "github.com/gocanto/blog/webkit/cli"
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

    defer (*logs).Close()
    defer (*dbConnection).Close()

    truncateDB(dbConnection, environment)

    seeder := seed.MakeSeeder(dbConnection)

    UserA, UserB := seeder.SeedUsers()

    seeder.SeedCategories()
    seeder.SeedTags()

    //seed.
    seeder.SeedPosts(UserA, UserB)

    // ------
    var posts database.Post
    if err := dbConnection.Sql().Model(&UserA).Association("Posts").Find(&posts); err != nil {
        panic(fmt.Sprintf("Could not find post by id: %v", err))
    } else {
        fmt.Println(posts)
    }

    //fmt.Println("--> ", UserA.ID, len(UserA.Posts))

    //cli.MakeTextColour("Done", cli.Green).Println()
}

func truncateDB(dbConnection *database.Connection, environment *env.Environment) {
    if environment.App.IsProduction() {
        panic(textColour.Print())
    }

    truncate := database.MakeTruncate(dbConnection, environment)

    if err := truncate.Execute(); err != nil {
        panic(err)
    }
}
