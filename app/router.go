package main

import (
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/logger"
	"github.com/gocanto/blog/app/support"
	"github.com/gocanto/blog/app/users"
	"net/http"
)

type Router struct {
	mux       *http.ServeMux
	env       *env.Environment
	container *Container
}

type Container struct {
	validator *support.Validator
	logs      *logger.Managers
	orm       *database.Orm
}

func makeRouter(mux *http.ServeMux, env *env.Environment, container *Container) Router {
	return Router{
		mux:       mux,
		env:       env,
		container: container,
	}
}

func (router Router) registerUsers() {
	provider := users.MakeProvider(
		users.MakeRepository(router.container.orm),
		router.container.validator,
	)

	provider.Register(router.mux)
}
