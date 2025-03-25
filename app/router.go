package main

import (
	"github.com/gocanto/blog/app/contracts"
	"github.com/gocanto/blog/app/support"
	"github.com/gocanto/blog/app/user"
	"net/http"
)

type Router struct {
	mux       *http.ServeMux
	validator *support.Validator
	logger    *contracts.LogsDriver
}

func getRouter(mux *http.ServeMux, logger *contracts.LogsDriver) Router {
	return Router{
		mux:       mux,
		validator: verifier,
		logger:    logger,
	}
}

func (r Router) registerUsers(db *contracts.DatabaseDriver) {
	users := user.RegisterProvider(
		user.NewRepository(db),
		r.validator,
	)

	users.Register(r.mux)
}
