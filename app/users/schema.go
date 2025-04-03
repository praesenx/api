package users

import (
	"github.com/gocanto/blog/app/kernel"
)

type HandleUsers struct {
	Validator  *kernel.Validator
	Repository *Repository
}

type CreatedUser struct {
	UUID string `json:"uuid"`
}
