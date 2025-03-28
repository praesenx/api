package users

import (
	"github.com/gocanto/blog/app/database"
	"github.com/google/uuid"
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

func (r Repository) Create(attr CreateRequestBag) (*CreatedUser, error) {
	user := &database.User{
		UUID:              uuid.New().String(),
		FirstName:         attr.FirstName,
		LastName:          attr.LastName,
		Username:          attr.Username,
		DisplayName:       attr.DisplayName,
		Email:             attr.Email,
		PasswordHash:      "asassas",
		Token:             "gocanto",
		Bio:               attr.Bio,
		ProfilePictureURL: attr.ProfilePictureURL,
	}

	orm := *r.Connection
	result := orm.Driver().Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &CreatedUser{}, nil
}
