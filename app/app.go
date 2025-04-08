package main

import (
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/users"
	"github.com/gocanto/blog/app/webkit"
	"github.com/gocanto/blog/app/webkit/llogs"
	"github.com/gocanto/blog/app/webkit/middleware"
	"net/http"
)

type App struct {
	Validator    *webkit.Validator    `validate:"required"`
	Logs         *llogs.Driver        `validate:"required"`
	dbConnection *database.Connection `validate:"required"`
	AdminUser    *users.AdminUser     `validate:"required"`
	Env          *env.Environment     `validate:"required"`
	Mux          *http.ServeMux       `validate:"required"`
	Sentry       *webkit.Sentry       `validate:"required"`
}

func MakeApp(mux *http.ServeMux, app *App) *App {
	app.Mux = mux

	return app
}

func (app App) RegisterUsers() {
	stack := middleware.MakeMiddlewareStack(app.Env, func(seed string) bool {
		return app.AdminUser.IsAllowed(seed)
	})

	handler := users.UserHandler{
		Repository: users.MakeRepository(app.dbConnection, app.AdminUser),
		Validator:  app.Validator,
	}

	app.Mux.HandleFunc("POST /users", webkit.CreateHandle(
		stack.Push(
			handler.Create,
			stack.AdminUser,
		),
	))
}
