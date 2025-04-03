package main

import (
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/logger"
	"github.com/gocanto/blog/app/middleware"
	"github.com/gocanto/blog/app/reponse"
	"github.com/gocanto/blog/app/support"
	"github.com/gocanto/blog/app/users"
	"net/http"
)

type App struct {
	Validator *support.Validator `validate:"required"`
	Logs      *logger.Managers   `validate:"required"`
	Orm       *database.Orm      `validate:"required"`
	AdminUser *users.AdminUser   `validate:"required"`
	Env       *env.Environment   `validate:"required"`
	Mux       *http.ServeMux     `validate:"required"`
}

func MakeApp(mux *http.ServeMux, app *App) *App {
	app.Mux = mux

	return app
}

func (app App) RegisterUsers() {
	stack := middleware.MakeStack(app.Env, app.AdminUser)

	handler := users.HandleUsers{
		Repository: users.MakeRepository(app.Orm, app.AdminUser),
		Validator:  app.Validator,
	}

	app.Mux.HandleFunc("POST /users", reponse.CreateHandle(
		stack.Push(
			handler.Create,
			//stack.Logging,
			//stack.AdminUser,
		),
	))
}
