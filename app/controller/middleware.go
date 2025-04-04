package controller

import (
	"github.com/gocanto/blog/app/env"
	"net/http"
)

func (s MiddlewareStack) Logging(next BaseController) BaseController {
	return func(w http.ResponseWriter, r *http.Request) *HttpError {
		println("Incoming request:", r.Method, r.URL.Path)

		err := next(w, r)

		if err != nil {
			println("Handler returned error:", err.Message)
		} else {
			println("Handler completed successfully")
		}

		return err
	}
}

func (s MiddlewareStack) AdminUser(next BaseController) BaseController {
	return func(w http.ResponseWriter, r *http.Request) *HttpError {
		salt := r.Header.Get(env.ApiKeyHeader)

		if s.isAdminUser(salt) {
			return next(w, r)
		}

		return Unauthorised("Unauthorized", nil)
	}
}

func (s MiddlewareStack) isAdminUser(seed string) bool {
	return s.userAdminResolver(seed)
}

func (s MiddlewareStack) Push(handler BaseController, middlewares ...Middleware) BaseController {
	// Apply middleware in reverse order, so the first middleware in the list is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
