package users

import (
	"github.com/gocanto/blog/app/webkit"
)

type UserHandler struct {
	Validator  *webkit.Validator
	Repository *Repository
}
