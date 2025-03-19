package database

import (
	"database/sql"
	"fmt"
)

type Connection struct {
	url string
	db  *sql.DB
}

func Make(url string) (Connection, error) {
	connection, err := sql.Open("postgres", url)

	if err != nil {
		return Connection{}, err
	}

	return Connection{
		url: url,
		db:  connection,
	}, nil
}

func (receiver *Connection) Close() error {
	if err := receiver.db.Close(); err != nil {
		return err
	}

	return nil
}

func (receiver *Connection) Ping() error {
	if err := receiver.db.Ping(); err != nil {
		return err
	}

	fmt.Println("Connected ....")

	return nil
}

func (receiver *Connection) GetDB() *sql.DB {
	return receiver.db
}
