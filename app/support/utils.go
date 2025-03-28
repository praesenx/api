package support

import (
	"log/slog"
	baseHttp "net/http"
)

func CloseRequestBody(r *baseHttp.Request) func() {
	return func() {
		if err := r.Body.Close(); err != nil {
			slog.Warn("Error closing request body", "error", err)
		}
	}
}
