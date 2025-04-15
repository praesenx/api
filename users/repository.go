package users

import (
    "errors"
    "fmt"
    "github.com/gocanto/blog/database"
    "github.com/gocanto/blog/webkit/orm"
    "github.com/google/uuid"
    "gorm.io/gorm"
    "strings"
    "time"
)

type Repository struct {
    connection *database.Connection
    admin      *AdminUser
}

func MakeRepository(model *database.Connection, admin *AdminUser) *Repository {
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

    result := r.connection.Sql().Create(&user)

    if result.Error != nil {
        return nil, result.Error
    }

    return &CreatedUser{
        UUID: user.UUID,
    }, nil
}

func (r Repository) FindByUserName(username string) *database.User {
    user := &database.User{}

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

func (r Repository) FindPosts(author database.User) (*[]database.Post, error) {
    var posts []database.Post

    err := r.connection.Sql().
        Model(&database.User{}).
        Preload("Posts", func(db *gorm.DB) *gorm.DB {
            return db.
                Where("author_id = ?", author.ID).
                Where("published_at IS NOT NULL").
                Where("published_at IS NOT NULL").
                Where("deleted_at IS NULL").
                Order("created_at desc")
        }).
        Find(&posts).
        Error

    if orm.IsNotFound(err) {
        return nil, errors.New(fmt.Sprintf("posts not found for author [%s]: %s", author.Username, err.Error()))
    }

    if orm.IsFoundButHasErrors(err) {
        return nil, errors.New(fmt.Sprintf("issue retrieving author's [%s] posts: %s", author.Username, err.Error()))
    }

    return &posts, nil
}
