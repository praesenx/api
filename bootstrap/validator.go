package bootstrap

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Driver *validator.Validate
}

func MakeValidator() Validator {
	return Validator{
		Driver: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (receiver Validator) Validate(params interface{}) map[string]interface{} {
	err := receiver.Driver.Struct(params)
	var issues = make(map[string]interface{})

	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		//slog.Error("Error happened in validate Struct. Err: %e", err)
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			//return []
		}

		var issues = make(map[string]interface{})

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				issues[e.Field()] = e.Error()
				//fmt.Println(e.Namespace())
				fmt.Println(e.Field())
				//fmt.Println(e.StructNamespace())
				//fmt.Println(e.StructField())
				fmt.Println(e.Tag())
				//fmt.Println(e.ActualTag())
				//fmt.Println(e.Kind())
				//fmt.Println(e.Type())
				//fmt.Println(e.Value())
				//fmt.Println(e.Param())
				//fmt.Println(e.Error())
				fmt.Println()
			}
		}
	} else {
		return issues
	}

	return issues
}
