package users

import "github.com/gocanto/blog/database"

type Repository struct {
	Connection database.Connection
}

func (r Repository) Create(user User) (error, User) {
	return nil, User{}
}
