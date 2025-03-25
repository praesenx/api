package users

import (
	"github.com/gocanto/blog/app/contracts"
)

type Repository struct {
	Connection *contracts.DatabaseDriver
}

func NewRepository(connection *contracts.DatabaseDriver) *Repository {
	return &Repository{
		Connection: connection,
	}
}

func (r Repository) Create(attr CreateUsersRequestBag) (error, CreatedUser) {
	return nil, CreatedUser{}
}
