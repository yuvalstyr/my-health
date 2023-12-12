package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *fiber.App {
	app := fiber.New(
		fiber.Config{
			AppName: "my-health",
		})
	app.Use(logger.New())
	return app
}
