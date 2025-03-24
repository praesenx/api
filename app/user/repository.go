package user

import "github.com/gocanto/blog/app/database"

type Repository struct {
	Connection database.Connection
}

func NewRepository(connection database.Connection) *Repository {
	return &Repository{
		Connection: connection,
	}
}

func (r Repository) Create(attr CreateUsersRequestBag) (error, CreatedUser) {
	return nil, CreatedUser{}
}
