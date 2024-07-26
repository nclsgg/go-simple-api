package middleware

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
		Message     string
	}

	XValidator struct {
		Validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

var Validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	errs := Validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
				Error:       true,
				Message: fmt.Sprintf(
					"[%s]: '%v' | Needs to implement '%s'",
					err.Field(),
					err.Value(),
					err.Tag(),
				),
			})
		}
	}

	return validationErrors
}
