package users

import (
	"github.com/gocanto/blog/app/support"
)

type ResponseBag struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type Handler struct {
	Validator  *support.Validator
	Repository *Repository
}

type CreatedUser struct{}
