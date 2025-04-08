package middleware

import (
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/webkit"
)

type MiddlewaresStack struct {
	env               *env.Environment
	middleware        []Middleware
	userAdminResolver func(seed string) bool
}

type Middleware func(webkit.BaseHandler) webkit.BaseHandler

func MakeMiddlewareStack(env *env.Environment, userAdminResolver func(seed string) bool) *MiddlewaresStack {
	return &MiddlewaresStack{
		env:               env,
		userAdminResolver: userAdminResolver,
		middleware:        []Middleware{},
	}
}
