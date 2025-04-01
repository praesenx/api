package middleware

import (
	"github.com/gocanto/blog/app/reponse"
	"net/http"
)

// Middleware type that accepts a BaseHandler and returns another BaseHandler.
type Middleware func(reponse.BaseHandler) reponse.BaseHandler

// ApplyMiddleware chains multiple middleware to a BaseHandler.
func ApplyMiddleware(handler reponse.BaseHandler, middlewares ...Middleware) reponse.BaseHandler {
	// Apply middleware in reverse order, so the first middleware in the list
	// is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}

// LoggingMiddleware Example Middleware: Logging Middleware
func LoggingMiddleware(next reponse.BaseHandler) reponse.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *reponse.ResponseError {
		// Perform actions before the handler
		println("Incoming request:", r.Method, r.URL.Path)

		// Call the next handler
		err := next(w, r)

		// Perform actions after the handler
		if err != nil {
			println("Handler returned error:", err.Message)
		} else {
			println("Handler completed successfully")
		}

		return err
	}
}

func AuthenticationMiddleware(next reponse.BaseHandler) reponse.BaseHandler {
	return func(w http.ResponseWriter, r *http.Request) *reponse.ResponseError {
		// Example: Check for an authentication header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer valid_token" {
			return reponse.MakeUnauthorized("Unauthorized", nil)
		}

		// Call the next handler if authentication is successful
		return next(w, r)
	}
}
