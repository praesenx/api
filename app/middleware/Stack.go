package middleware

import (
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/response"
	"github.com/gocanto/blog/app/users"
)

type Stack struct {
	env        *env.Environment
	middleware []Middleware
	adminUser  *users.AdminUser
}

type Middleware func(response.BaseHandler) response.BaseHandler

func MakeStack(env *env.Environment, adminUser *users.AdminUser) *Stack {
	return &Stack{
		env:        env,
		adminUser:  adminUser,
		middleware: []Middleware{},
	}
}

func (s Stack) isAdminUser(seed string) bool {
	return s.adminUser.IsAllowed(seed)
}

func (s Stack) Push(handler response.BaseHandler, middlewares ...Middleware) response.BaseHandler {
	// Apply middleware in reverse order, so the first middleware in the list is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
