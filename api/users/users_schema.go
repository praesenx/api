package users

import (
	"github.com/gocanto/blog/webkit"
)

type UserHandler struct {
	Validator  *webkit.Validator
	Repository *Repository
}
