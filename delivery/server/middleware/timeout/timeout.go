package timeout

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func New(opt ...Option) fiber.Handler {
	cfg := ConfigDefault

	for _, o := range opt {
		o(&cfg)
	}

	return func(c *fiber.Ctx) error {
		userCtx := c.UserContext()
		userCtx, cancel := context.WithTimeout(userCtx, cfg.Duration)
		defer cancel()

		c.SetUserContext(userCtx)

		return c.Next()
	}
}
