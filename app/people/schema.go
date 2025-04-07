package people

import (
	"github.com/gocanto/blog/app/proxy"
)

type UsersHandler struct {
	Validator  *proxy.Validator
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
