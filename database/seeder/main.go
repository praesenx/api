package main

import (
    "fmt"
    "github.com/gocanto/blog/boost"
    "github.com/gocanto/blog/database"
    "github.com/gocanto/blog/database/seeder/seeds"
    "github.com/gocanto/blog/env"
    "github.com/gocanto/blog/pkg/cli"
    "os"
    "os/exec"
    "sync"
    "time"
)

var environment *env.Environment

func init() {
    secrets, _ := boost.Spark("./.env")

    environment = secrets
}

func main() {
    clearScreen()
    dbConnection := boost.MakeDbConnection(environment)
    logs := boost.MakeLogs(environment)

    defer (*logs).Close()
    defer (*dbConnection).Close()

    // [1] --- Create the Seeder Runner.
    seeder := seeds.MakeSeeder(dbConnection, environment)

    // [2] --- Truncate the DB.
    if err := seeder.TruncateDB(); err != nil {
        panic(err)
    } else {
        cli.MakeTextColour("DB Truncated successfully ...", cli.Green).Print()
        time.Sleep(2 * time.Second)
    }

    // [3] --- Seed users and posts sequentially because the below seeders depend on them.
    UserA, UserB := seeder.SeedUsers()
    posts := seeder.SeedPosts(UserA, UserB)

    categoriesChan := make(chan []database.Category)
    tagsChan := make(chan []database.Tag)

    go func() {
        defer close(categoriesChan)

        cli.MakeTextColour("Seeding categories ...", cli.Yellow).Print()
        categoriesChan <- seeder.SeedCategories()
    }()

    go func() {
        defer close(tagsChan)

        cli.MakeTextColour("Seeding tags ...", cli.Magenta).Print()
        tagsChan <- seeder.SeedTags()
    }()

    // [4] Use channels to concurrently seed categories and tags since they are main dependencies.
    categories := <-categoriesChan
    tags := <-tagsChan

    // [5] Use a WaitGroup to run independent seeding tasks concurrently.
    var wg sync.WaitGroup
    wg.Add(6)

    go func() {
        defer wg.Done()

        cli.MakeTextColour("Seeding comments ...", cli.Blue).Print()
        seeder.SeedComments(posts...)
    }()

    go func() {
        defer wg.Done()

        cli.MakeTextColour("Seeding likes ...", cli.Cyan).Print()
        seeder.SeedLikes(posts...)
    }()

    go func() {
        defer wg.Done()

        cli.MakeTextColour("Seeding posts-categories ...", cli.Gray).Print()
        seeder.SeedPostsCategories(categories, posts)
    }()

    go func() {
        defer wg.Done()

        cli.MakeTextColour("Seeding posts-tags ...", cli.Magenta).Print()
        seeder.SeedPostTags(tags, posts)
    }()

    go func() {
        defer wg.Done()

        cli.MakeTextColour("Seeding views ...", cli.Yellow).Print()
        seeder.SeedPostViews(posts, UserA, UserB)
    }()

    go func() {
        defer wg.Done()

        cli.MakeTextColour("Seeding Newsletters ...", cli.Green).Print()
        if err := seeder.SeedNewsLetters(); err != nil {
            cli.MakeTextColour(err.Error(), cli.Red).Print()
        }
    }()

    wg.Wait()

    cli.MakeTextColour("DB seeded as expected.", cli.Green).Print()
}

func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        message := fmt.Sprintf("Could not clear screen. Error: %s", err.Error())

        cli.MakeTextColour(message, cli.Red).Print()
    }
}
