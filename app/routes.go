package main

import (
	"github.com/gocanto/blog/app/user"
	"net/http"
)

func usersRoutes(mux *http.ServeMux) {
	repository := user.NewRepository(
		dbConnection,
	)

	service := user.NewProvider(repository, validator)
	_ = service.Register(mux)
}
