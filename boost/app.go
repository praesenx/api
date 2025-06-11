package boost

import (
	"github.com/oullin/api/api/users"
	"github.com/oullin/api/database"
	"github.com/oullin/api/env"
	"github.com/oullin/api/pkg"
	"github.com/oullin/api/pkg/llogs"
	"github.com/oullin/api/pkg/middleware"
	"net/http"
)

type App struct {
	Validator    *pkg.Validator       `validate:"required"`
	Logs         *llogs.Driver        `validate:"required"`
	DbConnection *database.Connection `validate:"required"`
	AdminUser    *users.AdminUser     `validate:"required"`
	Env          *env.Environment     `validate:"required"`
	Mux          *http.ServeMux       `validate:"required"`
	Sentry       *pkg.Sentry          `validate:"required"`
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

	app.Mux.HandleFunc("POST /users", pkg.CreateHandle(
		stack.Push(
			handler.Create,
			stack.AdminUser,
		),
	))
}
