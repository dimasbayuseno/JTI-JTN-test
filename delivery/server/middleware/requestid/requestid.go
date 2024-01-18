package requestid

import (
	"context"
	"encoding/json"
	"initial/infrastructure/shared"
	"initial/infrastructure/shared/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		uuidStr := uuid.New().String()

		userCtx := c.UserContext()
		userCtx = context.WithValue(userCtx, shared.ContextKeyRequestID, uuidStr)
		c.SetUserContext(userCtx)

		err := c.Next()
		// get respnse
		var response response.GeneralResponse
		errUnmarshal := json.Unmarshal(c.Response().Body(), &response)
		if errUnmarshal != nil {
			return err
		}

		// set request id
		response.RequestID = uuidStr
		resByte, errMarshal := json.Marshal(&response)
		if errMarshal != nil {
			return err
		}
		// set response
		c.Response().SetBody(resByte)

		return err
	}
}
