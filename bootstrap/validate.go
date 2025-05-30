package bootstrap

import (
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/gocanto/blog/pkg"
)

func GetDefaultValidate() *pkg.Validator {
	return pkg.MakeValidatorFrom(baseValidator.New(
		baseValidator.WithRequiredStructEnabled(),
	))
}
