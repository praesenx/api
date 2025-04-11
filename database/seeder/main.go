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
	if err := truncateDB(dbConnection, environment); err != nil {
		panic(err)
	}

	createUsers(dbConnection)

	fmt.Println("\nDone ...")
}

func truncateDB(dbConnection *database.Connection, environment *env.Environment) error {
	text, err := cli.MakePaddedTextColour("Senders are forbidden in production.", cli.Red)

	if err != nil {
		panic(err)
	}

	if environment.App.IsProduction() {
		fmt.Print(text.Get())
		os.Exit(1)
	}

	truncate := database.MakeTruncate(dbConnection, environment)

	return truncate.Execute()
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
