package validation

import (
	"fmt"
	"reflect"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

var messageTmpls = map[string]interface{}{
	"required":       "The %s field is required.",
	"requiredWith":   "The %s field is required when %s is present.",
	"noDuplications": "The %s must not have duplicated values.",
	"uuid":           "The %s field must be in uuid format.",
	"email":          "The %s field must be in email format.",
	"url":            "The %s field must be in url format.",
	"gt":             "The %s field must be greater than %s",
	"numeric":        "The %s field must be a number.",
	"number":         "The %s field must be an integer.",
	"min": map[string]string{
		"numeric": "The %s field must be at least %s.",
		"string":  "The %s field must be at least %s characters.",
		"array":   "The %s field must have at least %s items.",
	},
	"max": map[string]string{
		"numeric": "The %s field must bot be greater than %s.",
		"string":  "The %s field must not be greater than %s characters.",
		"array":   "The %s field must not have more than %s items.",
	},
}

// ParseValidationErr extracts error messages from validation error
func ParseValidationErr(err error) map[string][]string {
	errMessages := make(map[string][]string)

	for _, fieldError := range err.(validator.ValidationErrors) {
		fieldName := customizeFieldName(fieldError.Namespace())
		errMessages[fieldName] = buildFieldValidationErrorMessage(fieldError, fieldName)
	}

	return errMessages
}

func customizeFieldName(nameSpace string) string {
	//eg: Package.platformPackageId -> platformPackageId
	firstIndex := strings.Index(nameSpace, ".")
	if firstIndex != -1 {
		nameSpace = nameSpace[firstIndex+1:]
	}
	//eg: items[0].price -> items.0.price
	nameSpace = strings.Replace(nameSpace, "[", ".", 1)
	return strings.Replace(nameSpace, "]", "", 1)
}

func buildFieldValidationErrorMessage(fieldError validator.FieldError, fieldName string) []string {
	messageTemplate := "Validation failed for %s"
	switch msgs := messageTmpls[fieldError.Tag()].(type) {
	case string:
		messageTemplate = msgs
	case map[string]string:
		switch fieldError.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
			messageTemplate = msgs["numeric"]
		case reflect.Slice, reflect.Array, reflect.Map:
			messageTemplate = msgs["array"]
		case reflect.String:
			messageTemplate = msgs["string"]
		}
	}

	args := []interface{}{fieldName}
	if fieldError.Param() != "" {
		args = append(args, fieldError.Param())
	}
	message := fmt.Sprintf(messageTemplate, args...)

	return []string{message}
}
