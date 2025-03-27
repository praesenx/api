package users

import (
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
	mux.HandleFunc("POST /users", provider.usersHandler.create)
}
