package exception

import (
	"encoding/json"
	"initial/infrastructure/logger"
	"initial/infrastructure/shared"
	"initial/infrastructure/shared/response"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	_, parentSpan := tracer.Start(c.UserContext(), "ErrorHandler")
	defer func() {
		logger.Log(c.UserContext()).Error().Msg(err.Error())
		parentSpan.RecordError(err)
		parentSpan.End()
	}()

	_, validationError := err.(shared.ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		return c.Status(fiber.StatusBadRequest).JSON(response.GeneralResponse{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	_, notFoundError := err.(shared.NotFoundError)
	if notFoundError {
		return c.Status(fiber.StatusNotFound).JSON(response.GeneralResponse{
			Code:    404,
			Message: "Not Found",
			Data:    err.Error(),
		})
	}

	_, unauthorizedError := err.(shared.UnauthorizedError)
	if unauthorizedError {
		return c.Status(fiber.StatusUnauthorized).JSON(response.GeneralResponse{
			Code:    401,
			Message: "Unauthorized",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(response.GeneralResponse{
		Code:    500,
		Message: "General Error",
		Data:    err.Error(),
	})
}
