package views

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// ComponentsHandler is a Go function that handles a page view.
// It takes a *fiber.Ctx parameter and returns an error.
func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
