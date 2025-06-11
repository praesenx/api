package middleware

import (
	"fmt"
	"github.com/oullin/env"
	"github.com/oullin/pkg"
	"github.com/oullin/pkg/response"
	"log/slog"
	"net/http"
)

func (s MiddlewaresStack) Logging(next pkg.BaseHandler) pkg.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *response.Response {
		slog.Info(fmt.Sprintf("Incoming request: [method:%s] [path:%s].", r.Method, r.URL.Path))

		err := next(w, r)

		if err != nil {
			slog.Error(fmt.Sprintf("Handler returned error: %s", err))
		}

		return err
	}
}

func (s MiddlewaresStack) AdminUser(next pkg.BaseHandler) pkg.BaseHandler {
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

func (s MiddlewaresStack) Push(handler pkg.BaseHandler, middlewares ...Middleware) pkg.BaseHandler {
	// Apply middleware in reverse order, so the first middleware in the list is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
