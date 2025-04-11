package database

import (
	"fmt"
	"github.com/gocanto/blog/env"
)

type Truncate struct {
	database *Connection
	env      *env.Environment
}

func MakeTruncate(db *Connection, env *env.Environment) *Truncate {
	return &Truncate{
		database: db,
		env:      env,
	}
}

func (t Truncate) Execute() error {
	for i := len(Tables) - 1; i >= 0; i-- {
		t.database.Sql().Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", Tables[i]))
		fmt.Println(fmt.Sprintf("Table [%s] sucessfully truncated.", Tables[i]))
	}

	return nil
}
