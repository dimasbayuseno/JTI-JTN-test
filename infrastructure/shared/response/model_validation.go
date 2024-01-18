package response

import (
	"encoding/json"
	"initial/infrastructure/shared"

	"github.com/go-playground/validator/v10"
)

func Validate(modelValidate interface{}) (err error) {
	validate := validator.New()
	err = validate.Struct(modelValidate)
	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "this field is " + err.Tag(),
			})
		}

		jsonMessage, errJson := json.Marshal(messages)
		if errJson != nil {
			return errJson
		}

		return &shared.ValidationError{
			Message: string(jsonMessage),
		}
	}

	return nil
}
