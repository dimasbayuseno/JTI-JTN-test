package timeout

import (
	"context"
	"initial/infrastructure/shared/response"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func sleepWithContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
		return ctx.Err()
	case <-timer.C:
	}
	return nil
}
func TestTimoutDeadlineExceeded(t *testing.T) {
	app := fiber.New()

	app.Use(New(WithDuration(3 * time.Second)))

	app.Get("/", func(c *fiber.Ctx) error {
		err := sleepWithContext(c.UserContext(), 8*time.Second)
		if err != nil {
			return err
		}
		return c.JSON(response.GeneralResponse{
			Code:    200,
			Message: "success",
		})
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil), 10000)

	assert.NoError(t, err)

	bodyByte, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, context.DeadlineExceeded.Error(), string(bodyByte))
}
