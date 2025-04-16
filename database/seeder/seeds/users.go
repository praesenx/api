package seeds

import (
	"fmt"
	"github.com/gocanto/blog/database"
	"github.com/gocanto/blog/users"
	"github.com/google/uuid"
	"strings"
	"time"
)

type UsersSeed struct {
	db *database.Connection
}

type UsersAttrs struct {
	Username string
	Name     string
	IsAdmin  bool
}

func MakeUsersSeed(db *database.Connection) *UsersSeed {
	return &UsersSeed{
		db: db,
	}
}

func (s UsersSeed) Create(attrs UsersAttrs) database.User {
	pass, _ := users.MakePassword("password")

	user := database.User{
		UUID:         uuid.NewString(),
		FirstName:    attrs.Name,
		LastName:     "Tester",
		Username:     attrs.Username,
		DisplayName:  fmt.Sprintf("%s User", attrs.Name),
		Email:        fmt.Sprintf("%s@test.com", strings.Trim(attrs.Username, " ")),
		PasswordHash: pass.GetHash(),
		PublicToken:  uuid.NewString(),
		IsAdmin:      attrs.IsAdmin,
		Bio:          "Nam hendrerit nulla ut cursus laoreet.",
		VerifiedAt:   time.Now(),
	}

	s.db.Sql().Create(&user)
	fmt.Println("Created: ", attrs.Name)

	return user
}
