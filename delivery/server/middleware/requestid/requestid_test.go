package requestid

import (
	"encoding/json"
	"fmt"
	"initial/infrastructure/shared/response"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRequestIDSuccess(t *testing.T) {
	app := fiber.New()

	app.Use(New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(response.GeneralResponse{
			Code:    200,
			Message: "success",
		})
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	bodyJson, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var responseBody response.GeneralResponse
	err = json.Unmarshal(bodyJson, &responseBody)
	assert.Nil(t, err)

	reqId, err := uuid.Parse(responseBody.RequestID)
	assert.Nil(t, err)
	fmt.Println(reqId)
	assert.NotEqual(t, uuid.Nil, reqId)

}

func TestRequestIDNotGenerated(t *testing.T) {
	app := fiber.New()

	app.Use(New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	_, err := app.Test(httptest.NewRequest("GET", "/", nil))
	assert.Nil(t, err)
}
