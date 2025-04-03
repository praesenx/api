package middleware

import (
	"github.com/gocanto/blog/app/kernel"
	"github.com/gocanto/blog/app/support"
	"net/http"
)

func (s Stack) Logging(next kernel.BaseHandler) kernel.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *kernel.HttpException {
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

func (s Stack) AdminUser(next kernel.BaseHandler) kernel.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *kernel.HttpException {
		salt := r.Header.Get(support.ApiKeyHeader)

		if s.isAdminUser(salt) {
			return next(w, r)
		}

		return kernel.MakeUnauthorisedException("Unauthorized", nil)
	}
}
