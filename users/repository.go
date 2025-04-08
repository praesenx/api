package users

import (
	"errors"
	database2 "github.com/gocanto/blog/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Repository struct {
	connection *database2.Connection
	admin      *AdminUser
}

func MakeRepository(model *database2.Connection, admin *AdminUser) *Repository {
	return &Repository{
		connection: model,
		admin:      admin,
	}
}

func (r Repository) Create(attr CreateRequestBag) (*CreatedUser, error) {
	password, err := MakePassword(attr.Password)

	if err != nil {
		return nil, err
	}

	user := &database2.User{
		UUID:              uuid.New().String(),
		FirstName:         attr.FirstName,
		LastName:          attr.LastName,
		Username:          attr.Username,
		DisplayName:       attr.DisplayName,
		Email:             attr.Email,
		PasswordHash:      password.GetHash(),
		PublicToken:       attr.PublicToken,
		Bio:               attr.Bio,
		PictureFileName:   attr.PictureFileName,
		ProfilePictureURL: attr.ProfilePictureURL,
		VerifiedAt:        time.Now(),
		IsAdmin:           strings.Trim(attr.Username, " ") == adminUserName,
	}

	result := r.connection.Sql().Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &CreatedUser{
		UUID: user.UUID,
	}, nil
}

func (r Repository) FindByUserName(username string) *database2.User {
	user := &database2.User{}

	result := r.connection.Sql().
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
