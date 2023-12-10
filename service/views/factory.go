package views

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"personal/health-app/service/templates"
)

// PageView is a Go function that handles a page view.
// It takes a *fiber.Ctx parameter and returns an error.
func PageView(c *fiber.Ctx) error {
	httpHandler := templ.Handler(templates.Page())
	return adaptor.HTTPHandler(httpHandler)(c)
}
