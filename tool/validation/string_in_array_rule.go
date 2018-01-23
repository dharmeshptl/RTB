package validation

import (
	"go_rtb/internal/tool/helper"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

func validateStringInArray(fl validator.FieldLevel) bool {
	inputValues := fl.Field().Slice(0, fl.Field().Len())
	validateValues := strings.Split(fl.Param(), ";")
	for i := 0; i < inputValues.Len(); i++ {
		if helper.DoesStringArrayContain(inputValues.Index(i).String(), validateValues) == false {
			return false
		}
	}

	return true
}
