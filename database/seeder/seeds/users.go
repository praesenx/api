package seeds

import (
    "fmt"
    "github.com/gocanto/blog/database"
    "github.com/gocanto/blog/users"
    "github.com/google/uuid"
    "strings"
    "time"
)

type CreateUsersAttrs struct {
    DB       *database.Connection
    Username string
    Name     string
    IsAdmin  bool
}

func CreateUser(attrs CreateUsersAttrs) database.User {
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

    attrs.DB.Sql().Create(&user)
    fmt.Println("Created: ", attrs.Name)

    return user
}
