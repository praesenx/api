package users

import (
	"github.com/gocanto/blog/database"
	"github.com/gocanto/blog/packages/db"
)

type Repository struct {
	Connection database.Connection
}

func (r Repository) Create(user db.User) (error, db.User) {
	return nil, db.User{}
}
