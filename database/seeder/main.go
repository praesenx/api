package main

import (
	"fmt"
	"github.com/gocanto/blog/bootstrap"
	"github.com/gocanto/blog/env"
	"github.com/gocanto/blog/webkit"
)

var environment *env.Environment
var validator *webkit.Validator

func init() {
	secrets, validate := bootstrap.Spark("./.env")

	environment = secrets
	validator = validate
}

func main() {
	fmt.Println((*environment).DB.Port)

	conn := bootstrap.MakeDbConnection(environment)
	logs := bootstrap.MakeLogs(environment)

	defer (*logs).Close()
	defer (*conn).Close()

	conn.Ping()
	fmt.Println("Seeder main ...")
}
