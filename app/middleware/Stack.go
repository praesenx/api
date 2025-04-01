package middleware

import (
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/reponse"
)

type Stack struct {
	env        *env.Environment
	middleware []Middleware
}

type Middleware func(reponse.BaseHandler) reponse.BaseHandler

func MakeStack(env *env.Environment) *Stack {
	return &Stack{
		env:        env,
		middleware: []Middleware{},
	}
}

func (s Stack) ShouldRejectAction(seed string) bool {
	return s.env.Admin.IsNotAllowed(seed)
}

func (s Stack) Push(handler reponse.BaseHandler, middlewares ...Middleware) reponse.BaseHandler {
	// Apply middleware in reverse order, so the first middleware in the list is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
