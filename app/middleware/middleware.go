package middleware

import (
	"github.com/gocanto/blog/app/reponse"
	"github.com/gocanto/blog/app/support"
	"net/http"
)

func (s Stack) Logging(next reponse.BaseHandler) reponse.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *reponse.ResponseError {
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

func (s Stack) AdminUser(next reponse.BaseHandler) reponse.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *reponse.ResponseError {
		salt := r.Header.Get(support.ApiKeyHeader)

		if s.isAdminUser(salt) {
			return next(w, r)
		}

		return reponse.MakeUnauthorized("Unauthorized", nil)
	}
}
