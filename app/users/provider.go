package users

import (
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/middleware"
	"github.com/gocanto/blog/app/reponse"
	"github.com/gocanto/blog/app/support"
	"net/http"
)

type Provider struct {
	Env          *env.Environment
	usersHandler *Handler
}

func MakeProvider(repository *Repository, validator *support.Validator, env *env.Environment) *Provider {
	return &Provider{
		Env: env,
		usersHandler: &Handler{
			Validator:  validator,
			Repository: repository,
		},
	}
}

func (provider *Provider) Register(mux *http.ServeMux) {
	stack := middleware.MakeStack(provider.Env)

	mux.HandleFunc("POST /users", reponse.CreateHandle(
		stack.Push(
			provider.usersHandler.create,
			stack.Logging,
			stack.Admin,
		),
	))
}
