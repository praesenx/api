package main

import (
	"github.com/gocanto/blog/database"
	"github.com/gocanto/blog/env"
	users2 "github.com/gocanto/blog/users"
	webkit2 "github.com/gocanto/blog/webkit"
	"github.com/gocanto/blog/webkit/llogs"
	"github.com/gocanto/blog/webkit/middleware"
	"net/http"
)

type App struct {
	Validator    *webkit2.Validator   `validate:"required"`
	Logs         *llogs.Driver        `validate:"required"`
	dbConnection *database.Connection `validate:"required"`
	AdminUser    *users2.AdminUser    `validate:"required"`
	Env          *env.Environment     `validate:"required"`
	Mux          *http.ServeMux       `validate:"required"`
	Sentry       *webkit2.Sentry      `validate:"required"`
}

func MakeApp(mux *http.ServeMux, app *App) *App {
	app.Mux = mux

	return app
}

func (app App) RegisterUsers() {
	stack := middleware.MakeMiddlewareStack(app.Env, func(seed string) bool {
		return app.AdminUser.IsAllowed(seed)
	})

	handler := users2.UserHandler{
		Repository: users2.MakeRepository(app.dbConnection, app.AdminUser),
		Validator:  app.Validator,
	}

	app.Mux.HandleFunc("POST /users", webkit2.CreateHandle(
		stack.Push(
			handler.Create,
			stack.AdminUser,
		),
	))
}
