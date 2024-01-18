package server

import (
	"github.com/gofiber/fiber/v2"
	"initial/infrastructure/jwt"
)

func routeGroupV1(app *fiber.App, handler handler) {
	v1 := app.Group("v1")
	internal := v1.Group("internal")
	internal.Use(jwt.ValidateTokenMiddleware)
	sample := internal.Group("/sample")
	{
		sample.Post("/sample-data", handler.sampleHandler.CreateData)
		sample.Get("/sample-data", handler.sampleHandler.GetAllData)
	}
}
