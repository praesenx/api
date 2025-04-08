package webkit

import (
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/gocanto/blog/env"
)

type Sentry struct {
	Handler *sentryhttp.Handler
	Env     *env.Environment
	Options *sentryhttp.Options
}
