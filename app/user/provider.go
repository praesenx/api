package user

import (
	"github.com/gocanto/blog/app/support"
	"net/http"
)

type Provider struct {
	Repository *Repository
	Validator  *support.Validator
}

func NewProvider(repository *Repository, validator *support.Validator) *Provider {
	return &Provider{
		Repository: repository,
		Validator:  validator,
	}
}

// Register users routes
func (p *Provider) Register(mux *http.ServeMux) *Controller {
	users := Controller{
		Validator:  p.Validator,
		Repository: p.Repository,
	}

	mux.HandleFunc("POST /users", users.Create)

	return &users
}
