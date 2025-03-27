package users

import (
	"github.com/gocanto/blog/app/database"
)

type Repository struct {
	Connection *database.Driver
	User       *database.User
}

func MakeRepository(connection *database.Driver) *Repository {
	return &Repository{
		Connection: connection,
	}
}

func (r Repository) Create(attr CreateUsersRequestBag) (error, CreatedUser) {
	return nil, CreatedUser{}
}
