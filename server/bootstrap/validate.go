package bootstrap

import (
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/gocanto/blog/server/webkit"
)

func GetDefaultValidate() *webkit.Validator {
	return webkit.MakeValidatorFrom(baseValidator.New(
		baseValidator.WithRequiredStructEnabled(),
	))
}
