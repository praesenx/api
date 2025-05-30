package bootstrap

import (
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/gocanto/blog/pkgs"
)

func GetDefaultValidate() *pkgs.Validator {
	return pkgs.MakeValidatorFrom(baseValidator.New(
		baseValidator.WithRequiredStructEnabled(),
	))
}
