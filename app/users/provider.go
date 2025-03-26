package users

import (
	"github.com/gocanto/blog/app/support"
	"net/http"
)

type Provider struct {
	repository   *Repository
	validator    *support.Validator
	usersHandler Handler
}

func RegisterProvider(repository *Repository, validator *support.Validator) *Provider {
	return &Provider{
		repository: repository,
		validator:  validator,
		usersHandler: Handler{
			Validator:  validator,
			Repository: repository,
		},
	}
}

// Register users routes
func (p *Provider) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", (*p).usersHandler.create)
}
