package fiberapp

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() *fiber.App {
	app := fiber.New(
		fiber.Config{
			AppName: "my-health",
		})
	app.Use(logger.New())
	return app
}
