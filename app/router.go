package main

import (
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/logger"
	"github.com/gocanto/blog/app/support"
	"github.com/gocanto/blog/app/users"
	"net/http"
)

type Router struct {
	mux       *http.ServeMux
	validator *support.Validator
	logger    *logger.Managers
}

func getRouter(mux *http.ServeMux, logger *logger.Managers) Router {
	return Router{
		mux:       mux,
		validator: verifier,
		logger:    logger,
	}
}

func (r Router) registerUsers(db *database.Driver) {
	provider := users.RegisterProvider(
		users.NewRepository(db),
		r.validator,
	)

	provider.Register(r.mux)
}
