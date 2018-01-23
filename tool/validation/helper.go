package validation

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("valid")
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	registerCustomValidateFunc(validate)

	return validate
}

func registerCustomValidateFunc(validate *validator.Validate) {
	validate.RegisterValidation("string_in_array", validateStringInArray)
}
