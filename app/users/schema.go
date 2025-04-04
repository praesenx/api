package users

import (
	"github.com/gocanto/blog/app/proxy"
)

type HandleUsers struct {
	Validator  *proxy.Validator
	Repository *Repository
}

type CreatedUser struct {
	UUID string `json:"uuid"`
}
