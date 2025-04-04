package kernel

import (
	"github.com/gocanto/blog/app/env"
	"net/http"
)

type Middleware func(BaseHandler) BaseHandler

func (s MiddlewareStack) Logging(next BaseHandler) BaseHandler {
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

func (s MiddlewareStack) AdminUser(next BaseHandler) BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *HttpError {
		salt := r.Header.Get(env.ApiKeyHeader)

		if s.isAdminUser(salt) {
			return next(w, r)
		}

		return Unauthorised("Unauthorized", nil)
	}
}
