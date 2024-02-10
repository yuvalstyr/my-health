package main

import (
	"log"
	"os"
	"personal/health-app/service/daos"
	"personal/health-app/service/datebase"
	"personal/health-app/service/handlers"
	"personal/health-app/service/model"
	"personal/health-app/service/views"
	"personal/health-app/service/views/dashboard"

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
	daoFactory := daos.NewDAOs(dbInstance.DB)

	app := echo.New()
	app.Static("/", "assets")

	handlersInstance := handlers.NewHandlersFactory(*daoFactory)
	app.POST(":id/increment", handlersInstance.Counter.Increment)
	app.POST(":id/decrement", handlersInstance.Counter.Decrement)
	app.POST(":id/sum", handlersInstance.Counter.Sum)

	app.GET("/", func(ctx echo.Context) error {
		var dishes []model.MealDish
		// TODO: change this to be dynamic, need to create meals for dishes
		res := dbInstance.DB.Find(&dishes)
		if res.Error != nil {
			return res.Error
		}

		activities, err := daoFactory.ActivityDAO.GetActivityDetails("", "2024-02-02")
		if err != nil {
			return errors.Wrap(err, "could not get activity details")
		}
		if len(activities) == 0 {
		}
		var activityTypes []model.ActivityType
		res = dbInstance.DB.Find(&activityTypes)
		if res.Error != nil {
			return res.Error
		}

		return views.Render(ctx, dashboard.Show(dishes, activities))
	})

	app.Start("localhost:4040")
}
