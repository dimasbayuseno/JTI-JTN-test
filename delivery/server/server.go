package server

import (
	"initial/configuration"
	"initial/delivery"
	"initial/delivery/server/middleware/requestid"
	"initial/delivery/server/middleware/timeout"
	"initial/exception"
	"initial/infrastructure/shared"
	"time"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.opentelemetry.io/otel/attribute"
)

func ServeHttp(container delivery.Container) *fiber.App {
	handler := SetupHandler(container)

	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(timeout.New(
		timeout.WithDuration(time.Duration(configuration.Env().GlobalTimeout) * time.Second),
	))

	if configuration.Env().Otel.EnableTracing {
		app.Use(otelfiber.Middleware(
			otelfiber.WithServerName(configuration.Env().ServicName),
			otelfiber.WithCustomAttributes(func(ctx *fiber.Ctx) []attribute.KeyValue {
				var attributes []attribute.KeyValue

				userCtx := ctx.UserContext()
				if reqId, ok := userCtx.Value(shared.ContextKeyRequestID).(string); ok {
					attributes = append(attributes, attribute.String(shared.ContextKeyRequestID.String(), reqId))
				}

				return attributes
			}),
		))
	}

	app.Use(recover.New())

	routeGroupV1(app, handler)

	return app
}
