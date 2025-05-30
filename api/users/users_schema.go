package users

import (
	"github.com/gocanto/blog/pkgs"
)

type UserHandler struct {
	Validator  *pkgs.Validator
	Repository *Repository
}
