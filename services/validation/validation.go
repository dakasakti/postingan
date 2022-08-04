package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type validation struct {
	validate *validator.Validate
}

func NewValidation() *validation {
	return &validation{validate: validator.New()}
}

func (v *validation) Validate(i interface{}) error {
	v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return v.validate.Struct(i)
}

func MessageValidate(err error) []string {
	var messages []string

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			messages = append(messages, fmt.Sprintf("%s is required", err.Field()))
		case "lowercase":
			messages = append(messages, fmt.Sprintf("%s must be lowercase", err.Field()))
		case "email":
			messages = append(messages, fmt.Sprintf("%s is not valid. example : username@domain.com", err.Field()))
		case "number":
			messages = append(messages, fmt.Sprintf("%s must be number", err.Field()))
		case "e164":
			messages = append(messages, fmt.Sprintf("%s is not valid. example : +6281234567890", err.Field()))
		case "min":
			messages = append(messages, fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param()))
		case "max":
			messages = append(messages, fmt.Sprintf("%s must be less than %s characters", err.Field(), err.Param()))
		}
	}

	return messages
}
