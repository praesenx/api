package middleware

import (
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/reponse"
	"github.com/gocanto/blog/app/users"
)

type Stack struct {
	env        *env.Environment
	middleware []Middleware
	adminUser  *users.AdminUser
}

type Middleware func(reponse.BaseHandler) reponse.BaseHandler

func MakeStack(env *env.Environment, adminUser *users.AdminUser) *Stack {
	return &Stack{
		env:        env,
		adminUser:  adminUser,
		middleware: []Middleware{},
	}
}

func (s Stack) AllowsAction(seed string) bool {
	return s.adminUser.IsAllowed(seed)
}

func (s Stack) Push(handler reponse.BaseHandler, middlewares ...Middleware) reponse.BaseHandler {
	// Apply middleware in reverse order, so the first middleware in the list is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
