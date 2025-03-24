package user

import "github.com/gocanto/blog/app/database"

type Repository struct {
	Connection database.Connection
}

func (r Repository) Create(attr CreateUsersRequestBag) (error, CreatedUser) {
	return nil, CreatedUser{}
}
