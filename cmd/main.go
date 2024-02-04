package main

import (
	"log"
	"os"
	"personal/health-app/service/datebase"
	"personal/health-app/service/handlers"
	"personal/health-app/service/model"
	"personal/health-app/service/templates"
	"personal/health-app/service/views"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")

	dbInstance, err := datebase.New(dbURL)
	if err != nil {
		println(errors.Wrapf(err, "failed to connect database").Error())
		panic(err)
	}

	app := echo.New()
	app.Static("../index.css", "styles")

	handlersInstance := handlers.NewHandlersFactory(dbInstance.DB)
	app.POST(":id/increment", handlersInstance.Counter.Increment)
	app.POST(":id/decrement", handlersInstance.Counter.Decrement)

	app.GET("/", func(ctx echo.Context) error {
		var dishes []model.MealDish
		// TODO: change this to be dynamic, need to create meals for dishes
		res := dbInstance.DB.Find(&dishes)
		if res.Error != nil {
			return res.Error
		}
		var activities []model.Activity
		res = dbInstance.
			DB.
			// TODO: change this to be dynamic, maybe remove from / page
			Where("activities.date = ?", "2024-02-02").
			Find(&activities)
		if res.Error != nil {
			return res.Error
		}
		var activityTypes []model.ActivityType
		res = dbInstance.DB.Find(&activityTypes)
		if res.Error != nil {
			return res.Error
		}

		return views.Render(ctx, templates.Page(dishes, activities))
	})

	app.Start("localhost:4040")
}
