package views

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// ComponentsHandler is a Go function that handles a page view.
// It takes a *fiber.Ctx parameter and returns an error.
func ComponentsHandler(c *fiber.Ctx, components templ.Component) error {
	httpHandler := templ.Handler(components)
	return adaptor.HTTPHandler(httpHandler)(c)
}
