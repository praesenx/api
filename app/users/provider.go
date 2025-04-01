package users

import (
	"github.com/gocanto/blog/app/middleware"
	"github.com/gocanto/blog/app/reponse"
	"github.com/gocanto/blog/app/support"
	"net/http"
)

type Provider struct {
	usersHandler *Handler
}

func MakeProvider(repository *Repository, validator *support.Validator) *Provider {
	return &Provider{
		usersHandler: &Handler{
			Validator:  validator,
			Repository: repository,
		},
	}
}

func (provider *Provider) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", reponse.CreateHandle(
		middleware.ApplyMiddleware(
			provider.usersHandler.create,
			middleware.LoggingMiddleware,
			middleware.AuthenticationMiddleware,
		),
	))

	//mux.HandleFunc("POST /users", reponse.CreateHandle(
	//	provider.usersHandler.create,
	//))
}
