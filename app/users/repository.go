package users

import (
	"errors"
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/env"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type Repository struct {
	model *database.Orm
	admin env.GlobalAdmin
}

func MakeRepository(model *database.Orm, admin env.GlobalAdmin) *Repository {
	return &Repository{
		model: model,
		admin: admin,
	}
}

func (r Repository) Create(attr CreateRequestBag) (*CreatedUser, error) {
	password, err := MakePassword(attr.Password)
	if err != nil {
		return nil, err
	}

	user := &database.User{
		UUID:              uuid.New().String(),
		FirstName:         attr.FirstName,
		LastName:          attr.LastName,
		Username:          attr.Username,
		DisplayName:       attr.DisplayName,
		Email:             attr.Email,
		PasswordHash:      password.GetHash(),
		Token:             r.admin.Token, // sha256
		TokenSalt:         r.admin.Salt,
		Bio:               attr.Bio,
		ProfilePictureURL: attr.ProfilePictureURL,
	}

	result := r.model.DB().Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &CreatedUser{
		UUID: user.UUID,
	}, nil
}

func (r Repository) FindByUserName(username string) *database.User {
	user := &database.User{}

	result := r.model.DB().
		Where("username = ?", username).
		First(&user)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	if strings.Trim(user.UUID, " ") != "" {
		return user
	}

	return nil
}
