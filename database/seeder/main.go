package main

import (
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

	// --- Tests
	//adminUser := bootstrap.MakeAdminUser(environment)
	//usersRepo := users.Repository{Connection: dbConnection, Admin: adminUser}
	//result, err := usersRepo.FindPosts(UserA)
	//fmt.Println("-> resul: ", len(result), " err: ", err)

	truncateDB(dbConnection, environment)

	seeder := seed.MakeSeeder(dbConnection)

	UserA, UserB := seeder.SeedUsers()
	posts := seeder.SeedPosts(UserA, UserB)

	seeder.SeedCategories()
	seeder.SeedTags()
	seeder.SeedComments(posts...)
	seeder.SeedLikes(posts...)

	cli.MakeTextColour("Done", cli.Green).Println()
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
