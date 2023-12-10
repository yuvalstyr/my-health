package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"personal/health-app/service/views"
)

func main() {
	app := fiber.New()
	app.Get("/", views.PageView)

	err := app.Listen(":4040")
	if err != nil {
		println(errors.Wrapf(err, "failed to start server").Error())
	}

}
