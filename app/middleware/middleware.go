package middleware

import (
	"github.com/gocanto/blog/app/reponse"
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

func (s Stack) Admin(next reponse.BaseHandler) reponse.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *reponse.ResponseError {
		salt := r.Header.Get("X-API-Key")

		if s.AllowsAction(salt) {
			return next(w, r)
		}

		return reponse.MakeUnauthorized("Unauthorized", nil)
	}
}
