package middleware

import (
	"github.com/gocanto/blog/env"
	"github.com/gocanto/blog/server/webkit"
	"github.com/gocanto/blog/server/webkit/response"
	"net/http"
)

func (s MiddlewaresStack) Logging(next webkit.BaseHandler) webkit.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *response.Response {
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

func (s MiddlewaresStack) AdminUser(next webkit.BaseHandler) webkit.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *response.Response {
		salt := r.Header.Get(env.ApiKeyHeader)

		if s.isAdminUser(salt) {
			return next(w, r)
		}

		return response.Unauthorized("Unauthorized", nil)
	}
}

func (s MiddlewaresStack) isAdminUser(seed string) bool {
	return s.userAdminResolver(seed)
}

func (s MiddlewaresStack) Push(handler webkit.BaseHandler, middlewares ...Middleware) webkit.BaseHandler {
	// Apply middleware in reverse order, so the first middleware in the list is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
