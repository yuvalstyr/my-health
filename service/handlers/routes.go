package handlers

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Echo, handlers *Factory) {
	app.POST(":id/increment", handlers.Counter.Increment)
	app.POST(":id/decrement", handlers.Counter.Decrement)
	app.POST(":id/sum", handlers.Counter.Sum)
	app.POST("/dish/add", handlers.DishHandler.Create)
	app.DELETE("/dish/:id", handlers.DishHandler.Delete)
	app.PUT("/dish/:id", handlers.DishHandler.Update)
	app.GET("/dish/:id/edit", handlers.DishHandler.Get)
}
