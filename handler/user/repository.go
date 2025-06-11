package user

import (
	"fmt"
	"github.com/gocanto/blog/database"
	"github.com/gocanto/blog/pkg/gorm"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Repository struct {
	Connection *database.Connection
	Admin      *AdminUser
}

func MakeRepository(model *database.Connection, admin *AdminUser) *Repository {
	return &Repository{
		Connection: model,
		Admin:      admin,
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
		PictureFileName:   attr.PictureFileName,
		ProfilePictureURL: attr.ProfilePictureURL,
		VerifiedAt:        time.Now(),
		IsAdmin:           strings.Trim(attr.Username, " ") == adminUserName,
	}

	result := r.Connection.Sql().Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &CreatedUser{
		UUID: user.UUID,
	}, nil
}

func (r Repository) FindByUserName(username string) *database.User {
	user := &database.User{}

	result := r.Connection.Sql().
		Where("username = ?", username).
		First(&user)

	if gorm.HasDbIssues(result.Error) {
		return nil
	}

	if strings.Trim(user.UUID, " ") != "" {
		return user
	}

	return nil
}

func (r Repository) FindPosts(author database.User) ([]database.Post, error) {
	var posts []database.Post

	err := r.Connection.Sql().
		Model(&database.Post{}).
		Where("author_id = ?", author.ID).
		Where("published_at IS NOT NULL").
		Where("deleted_at IS NULL").
		Order("created_at desc").
		Find(&posts).
		Error

	if gorm.IsNotFound(err) {
		return nil, fmt.Errorf("posts not found for author [%s]: %s", author.Username, err.Error())
	}

	if gorm.IsFoundButHasErrors(err) {
		return nil, fmt.Errorf("issue retrieving author's [%s] posts: %s", author.Username, err.Error())
	}

	return posts, nil
}
