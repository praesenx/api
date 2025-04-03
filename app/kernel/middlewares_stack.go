package kernel

import (
	"github.com/gocanto/blog/app/env"
)

type MiddlewareStack struct {
	env               *env.Environment
	middleware        []Middleware
	userAdminResolver func(seed string) bool
}

func MakeMiddlewareStack(env *env.Environment, userAdminResolver func(seed string) bool) *MiddlewareStack {
	return &MiddlewareStack{
		env:               env,
		userAdminResolver: userAdminResolver,
		middleware:        []Middleware{},
	}
}

func (s MiddlewareStack) isAdminUser(seed string) bool {
	return s.userAdminResolver(seed)
}

func (s MiddlewareStack) Push(handler BaseHandler, middlewares ...Middleware) BaseHandler {
	// Apply middleware in reverse order, so the first middleware in the list is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
