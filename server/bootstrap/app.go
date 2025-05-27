package bootstrap

import (
	"github.com/gocanto/blog/api/users"
	"github.com/gocanto/blog/database"
	"github.com/gocanto/blog/env"
	webkit2 "github.com/gocanto/blog/server/webkit"
	"github.com/gocanto/blog/server/webkit/llogs"
	"github.com/gocanto/blog/server/webkit/middleware"
	"net/http"
)

type App struct {
	Validator    *webkit2.Validator   `validate:"required"`
	Logs         *llogs.Driver        `validate:"required"`
	DbConnection *database.Connection `validate:"required"`
	AdminUser    *users.AdminUser     `validate:"required"`
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

	handler := users.UserHandler{
		Repository: users.MakeRepository(app.DbConnection, app.AdminUser),
		Validator:  app.Validator,
	}

	app.Mux.HandleFunc("POST /users", webkit2.CreateHandle(
		stack.Push(
			handler.Create,
			stack.AdminUser,
		),
	))
}
