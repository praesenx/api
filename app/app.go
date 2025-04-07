package main

import (
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/logs"
	"github.com/gocanto/blog/app/people"
	"github.com/gocanto/blog/app/proxy"
	"github.com/gocanto/blog/app/webkit"
	"github.com/gocanto/blog/app/webkit/middleware"
	"net/http"
)

type App struct {
	Validator *proxy.Validator  `validate:"required"`
	Logs      *logs.Driver      `validate:"required"`
	Orm       *database.Orm     `validate:"required"`
	AdminUser *people.AdminUser `validate:"required"`
	Env       *env.Environment  `validate:"required"`
	Mux       *http.ServeMux    `validate:"required"`
}

func MakeApp(mux *http.ServeMux, app *App) *App {
	app.Mux = mux

	return app
}

func (app App) RegisterUsers() {
	stack := middleware.MakeMiddlewareStack(app.Env, func(seed string) bool {
		return app.AdminUser.IsAllowed(seed)
	})

	handler := people.UsersHandler{
		Repository: people.MakeRepository(app.Orm, app.AdminUser),
		Validator:  app.Validator,
	}

	app.Mux.HandleFunc("POST /users", webkit.CreateHandle(
		stack.Push(
			handler.Create,
			stack.AdminUser,
		),
	))
}
