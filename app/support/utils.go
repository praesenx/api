package support

import (
	"log/slog"
	"net/http"
)

func CloseRequestBody(request *http.Request) func() {
	return func() {
		if err := request.Body.Close(); err != nil {
			slog.Warn("Error closing request body", "error", err)
		}
	}
}
