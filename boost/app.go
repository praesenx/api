package boost

import (
	"github.com/oullin/database"
	"github.com/oullin/env"
	"github.com/oullin/handler/user"
	"github.com/oullin/pkg"
	"github.com/oullin/pkg/llogs"
	"github.com/oullin/pkg/middleware"
	"net/http"
)

type App struct {
	Validator    *pkg.Validator       `validate:"required"`
	Logs         *llogs.Driver        `validate:"required"`
	DbConnection *database.Connection `validate:"required"`
	AdminUser    *user.AdminUser      `validate:"required"`
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

	handler := user.RequestHandler{
		Repository: user.MakeRepository(app.DbConnection, app.AdminUser),
		Validator:  app.Validator,
	}

	app.Mux.HandleFunc("POST /users", pkg.CreateHandle(
		stack.Push(
			handler.Create,
			stack.AdminUser,
		),
	))
}
