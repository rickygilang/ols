package helpers

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]interface{} {
	errors := make(map[string]interface{})

	if reflect.TypeOf(err) == reflect.TypeOf((*validator.ValidationErrors)(nil)).Elem() {
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := ""

			switch e.Tag() {
			case "required":
				errMessage = fmt.Sprintf("%v is required", e.Field())
			case "max":
				errMessage = fmt.Sprintf("%v cannot be longer than %v", e.Field(), e.Param())
			case "min":
				errMessage = fmt.Sprintf("%v must be longer than %v", e.Field(), e.Param())
			case "email":
				errMessage = fmt.Sprintf("%v invalid email format", e.Field())
			case "len":
				errMessage = fmt.Sprintf("%v must be %v characters long", e.Field(), e.Param())
			case "numeric":
				errMessage = fmt.Sprintf("%v must be numeric", e.Field())
			default:
				errMessage = e.Error()
			}
			errors[e.Field()] = errMessage
		}
	} else {
		errors["other_error"] = err.Error()
	}

	return errors
}

// To dereference value of pointer string, if it is nil then return empty string
func DerefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
