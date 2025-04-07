package users

import (
	"github.com/gocanto/blog/app/webkit"
)

type UsersHandler struct {
	Validator  *webkit.Validator
	Repository *Repository
}

type CreatedUser struct {
	UUID string `json:"uuid"`
}

type UserProfilePhoto struct {
	file       []byte
	payload    []byte
	headerName string
}
