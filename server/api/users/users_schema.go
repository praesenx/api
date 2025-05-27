package users

import (
	"github.com/gocanto/blog/server/webkit"
)

type UserHandler struct {
	Validator  *webkit.Validator
	Repository *Repository
}
