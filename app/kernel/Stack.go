package kernel

import (
	"github.com/gocanto/blog/app/env"
)

type Stack struct {
	env               *env.Environment
	middleware        []Middleware
	userAdminResolver func(seed string) bool
}

type Middleware func(BaseHandler) BaseHandler

func MakeStack(env *env.Environment, userAdminResolver func(seed string) bool) *Stack {
	return &Stack{
		env:               env,
		userAdminResolver: userAdminResolver,
		middleware:        []Middleware{},
	}
}

func (s Stack) isAdminUser(seed string) bool {
	return s.userAdminResolver(seed)
}

func (s Stack) Push(handler BaseHandler, middlewares ...Middleware) BaseHandler {
	// Apply middleware in reverse order, so the first middleware in the list is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
