package middleware

import (
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/webkit"
	"github.com/gocanto/blog/app/webkit/response"
	"net/http"
)

func (s MiddlewareStack) Logging(next webkit.BaseHandler) webkit.BaseHandler {
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

func (s MiddlewareStack) AdminUser(next webkit.BaseHandler) webkit.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *response.Response {
		salt := r.Header.Get(env.ApiKeyHeader)

		if s.isAdminUser(salt) {
			return next(w, r)
		}

		return response.Unauthorized("Unauthorized", nil)
	}
}

func (s MiddlewareStack) isAdminUser(seed string) bool {
	return s.userAdminResolver(seed)
}

func (s MiddlewareStack) Push(handler webkit.BaseHandler, middlewares ...Middleware) webkit.BaseHandler {
	// Apply middleware in reverse order, so the first middleware in the list is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
