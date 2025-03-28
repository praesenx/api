package users

import "github.com/gocanto/blog/app/support"

type Handler struct {
	Validator  *support.Validator
	Repository *Repository
}

type CreatedUser struct {
	UUID string `json:"uuid"`
}
