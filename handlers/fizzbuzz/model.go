package fizzbuzz

import (
	"fizz-buzz-api/services/fizzbuzz"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Int1  int    `json:"int1" validate:"required,gt=0"`
	Int2  int    `json:"int2" validate:"required,gt=0"`
	Limit int    `json:"limit" validate:"required,gt=0"`
	Str1  string `json:"str1" validate:"required,min=1"`
	Str2  string `json:"str2" validate:"required,min=1"`
}

// Map of custom error messages
var customErrorMessages = map[string]string{
	"Int1.required":  "int1 is required and must be greater than zero",
	"Int1.gt":        "int1 must be greater than zero",
	"Int2.required":  "int2 is required and must be greater than zero",
	"Int2.gt":        "int2 must be greater than zero",
	"Limit.required": "limit is required and must be greater than zero",
	"Limit.gt":       "limit must be greater than zero",
	"Str1.required":  "str1 is required and must not be empty",
	"Str1.min":       "str1 must not be empty",
	"Str2.required":  "str2 is required and must not be empty",
	"Str2.min":       "str2 must not be empty",
}

// Validate validates a FizzBuzzRequest
func (req Request) Validate() (validationErrors []string) {
	v := validator.New()
	if err := v.Struct(req); err != nil {
		// Collect validation errors
		for _, err := range err.(validator.ValidationErrors) {
			// Generate a unique key for each field + tag, e.g., "Int1.required"
			key := fmt.Sprintf("%s.%s", err.Field(), err.Tag())
			if customMessage, exists := customErrorMessages[key]; exists {
				validationErrors = append(validationErrors, customMessage)
			} else {
				// Default error message if not found in custom messages
				validationErrors = append(validationErrors, err.Error())
			}
		}
	}
	return
}
func (req Request) ToCommand() fizzbuzz.Command {
	return fizzbuzz.Command{
		Int1:  req.Int1,
		Int2:  req.Int2,
		Limit: req.Limit,
		Str1:  req.Str1,
		Str2:  req.Str2,
	}
}
