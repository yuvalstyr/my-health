package controller

import (
	"personal/health-app/service/handlers"

	"github.com/labstack/echo/v4"
)

func NewControllers(handlers handlers.Factory) []echo.Route {
	return []echo.Route{
		{},
	}
}
