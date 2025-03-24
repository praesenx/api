package user

import (
	"github.com/gocanto/blog/app/support"
)

type CreateUsersRequestBag struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type ResponseBag struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type Controller struct {
	Validator  *support.Validator
	Repository *Repository
}

type CreatedUser struct{}
