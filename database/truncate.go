package database

import "fmt"

type DbTruncate struct {
	Db    *Connection
	Table string
}

func MakeTruncate(db *Connection, table string) *DbTruncate {
	return &DbTruncate{
		Db:    db,
		Table: table,
	}
}

func (t DbTruncate) Truncate() {
	t.Db.Sql().Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", t.Table))
}

func (t DbTruncate) Fresh() {
	//db.Exec("DROP TABLE users")
	//t.Db.Sql().Exec(fmt.Sprintf("TRUNCATE TABLE %s", t.Table))
}
