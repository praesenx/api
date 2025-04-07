package people

import (
	"errors"
	"github.com/gocanto/blog/app/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Repository struct {
	model *database.Orm
	admin *AdminUser
}

func MakeRepository(model *database.Orm, admin *AdminUser) *Repository {
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
		PublicToken:       attr.PublicToken,
		Bio:               attr.Bio,
		ProfilePictureURL: attr.ProfilePictureURL,
		VerifiedAt:        time.Now(),
		IsAdmin:           strings.Trim(attr.Username, " ") == adminUserName,
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
