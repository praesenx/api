package middleware

import (
	"github.com/gocanto/blog/app/response"
	"github.com/gocanto/blog/app/support"
	"net/http"
)

func (s Stack) Logging(next response.BaseHandler) response.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *response.HttpException {
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

func (s Stack) AdminUser(next response.BaseHandler) response.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *response.HttpException {
		salt := r.Header.Get(support.ApiKeyHeader)

		if s.isAdminUser(salt) {
			return next(w, r)
		}

		return response.MakeUnauthorisedException("Unauthorized", nil)
	}
}
