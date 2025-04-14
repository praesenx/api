package main

import (
    "fmt"
    "github.com/gocanto/blog/bootstrap"
    "github.com/gocanto/blog/database"
    "github.com/gocanto/blog/database/seeder/seed"
    "github.com/gocanto/blog/env"
    "github.com/gocanto/blog/webkit/cli"
    "github.com/google/uuid"
    "time"
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
    textColour = makeTextColour("Done", cli.Green)

    defer (*logs).Close()
    defer (*dbConnection).Close()

    truncateDB(dbConnection, environment)

    users := seed.MakeUsersSeed(dbConnection)
    UserA := users.Create(seed.UsersAttrs{
        Username: "gocanto",
        Name:     "Gus",
        IsAdmin:  true,
    })
    UserB := users.Create(seed.UsersAttrs{
        Username: "li",
        Name:     "liane",
        IsAdmin:  false,
    })

    // ------
    posts := seed.MakePostsSeed(dbConnection)
    posts.CreatePosts(seed.PostsAttrs{
        AuthorID:    UserA.ID,
        Slug:        fmt.Sprintf("post-slug-%s", uuid.NewString()),
        Title:       fmt.Sprintf("Post %s title", uuid.NewString()),
        Excerpt:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris", uuid.NewString()),
        Content:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris. Nullam tempor velit id mi luctus, a scelerisque libero accumsan. In hac habitasse platea dictumst. Cras ac nunc nec massa tristique fringilla.", uuid.NewString()),
        PublishedAt: time.Now(),
        Author:      UserA,
        Categories:  []database.Category{},
        Tags:        []database.Tag{},
        PostViews:   []database.PostView{},
        Comments:    []database.Comment{},
        Likes:       []database.Like{},
    }, 1)

    posts.CreatePosts(seed.PostsAttrs{
        AuthorID:    UserB.ID,
        Slug:        fmt.Sprintf("post-slug-%s", uuid.NewString()),
        Title:       fmt.Sprintf("Post %s title", uuid.NewString()),
        Excerpt:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris", uuid.NewString()),
        Content:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris. Nullam tempor velit id mi luctus, a scelerisque libero accumsan. In hac habitasse platea dictumst. Cras ac nunc nec massa tristique fringilla.", uuid.NewString()),
        PublishedAt: time.Now(),
        Author:      UserB,
        Categories:  []database.Category{},
        Tags:        []database.Tag{},
        PostViews:   []database.PostView{},
        Comments:    []database.Comment{},
        Likes:       []database.Like{},
    }, 1)

    fmt.Println(textColour.Get())
}

func makeTextColour(message, colour string) *cli.TextColour {
    textColour, err := cli.MakePaddedTextColour(message, colour)

    if err != nil {
        panic(err)
    }

    return textColour
}

func truncateDB(dbConnection *database.Connection, environment *env.Environment) {
    if environment.App.IsProduction() {
        panic(textColour.Get())
    }

    truncate := database.MakeTruncate(dbConnection, environment)

    if err := truncate.Execute(); err != nil {
        panic(err)
    }
}
