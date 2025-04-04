package main

import (
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/kernel"
	"github.com/gocanto/blog/app/logs"
	"github.com/gocanto/blog/app/proxy"
	"github.com/gocanto/blog/app/users"
	"net/http"
)

type App struct {
	Validator *proxy.Validator `validate:"required"`
	Logs      *logs.Driver     `validate:"required"`
	Orm       *database.Orm    `validate:"required"`
	AdminUser *users.AdminUser `validate:"required"`
	Env       *env.Environment `validate:"required"`
	Mux       *http.ServeMux   `validate:"required"`
}

func MakeApp(mux *http.ServeMux, app *App) *App {
	app.Mux = mux

	return app
}

func (app App) RegisterUsers() {
	stack := kernel.MakeMiddlewareStack(app.Env, func(seed string) bool {
		return app.AdminUser.IsAllowed(seed)
	})

	handler := users.HandleUsers{
		Repository: users.MakeRepository(app.Orm, app.AdminUser),
		Validator:  app.Validator,
	}

	app.Mux.HandleFunc("POST /users", kernel.CreateHandle(
		stack.Push(
			handler.Create,
			stack.Logging,
			stack.AdminUser,
		),
	))
}
