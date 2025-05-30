package users

import (
	"github.com/gocanto/blog/pkg"
)

type UserHandler struct {
	Validator  *pkg.Validator
	Repository *Repository
}
