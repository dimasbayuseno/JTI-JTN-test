package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}
