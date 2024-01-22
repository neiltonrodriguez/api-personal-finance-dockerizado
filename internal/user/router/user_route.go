package router

import (
	"personal-finance-api/domain"
	"personal-finance-api/internal/user/handler"
	"personal-finance-api/middlewares"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(
	app *fiber.App,

) {
	config := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(domain.GlobalConfig.JwtSecretKey)}})

	route := app.Group(
		"v1/user",
		middlewares.Restricted,
	)

	route.Get("/", func(fiberCtx *fiber.Ctx) error {
		return handler.Get(fiberCtx)
	})

	route.Get("/:id", func(fiberCtx *fiber.Ctx) error {
		return handler.GetById(fiberCtx)
	})

	route.Put("/:id", func(fiberCtx *fiber.Ctx) error {
		return handler.Update(fiberCtx)
	})

	route.Post("/", func(fiberCtx *fiber.Ctx) error {
		return handler.Create(fiberCtx)
	})

	route.Use(config)

	route.Delete("/:id", func(fiberCtx *fiber.Ctx) error {
		return handler.Delete(fiberCtx)
	})

}