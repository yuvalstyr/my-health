package main

import (
	"fmt"
	"log"
	"os"
	"personal/health-app/service/daos"
	"personal/health-app/service/datebase"
	"personal/health-app/service/handlers"
	"personal/health-app/service/model"
	"personal/health-app/service/templates"
	"personal/health-app/service/views"
	"personal/health-app/service/views/components"
	"strconv"
	"time"

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
	app.Static("../index.css", "styles")

	handlersInstance := handlers.NewHandlersFactory(*daoFactory)
	app.POST(":id/increment", handlersInstance.Counter.Increment)
	app.POST(":id/decrement", handlersInstance.Counter.Decrement)
	app.POST(":id/sum", func(ctx echo.Context) error {
		sumID := ctx.Param("id")
		date := ctx.QueryParam("date")
		dateParsed, err := time.Parse("2006-01-02", date)
		if err != nil {
			return errors.Wrap(err, "invalid date")
		}
		value := ctx.FormValue("sum")
		valueParsed, err := strconv.Atoi(value)
		if err != nil {
			return errors.Wrap(err, "invalid value")
		}
		activities, err := daoFactory.ActivityDAO.GetActivityDetails(sumID, date)
		if err != nil {
			return errors.Wrap(err, "failed to get activity details")
		}
		activity := &model.Activity{
			ID:     sumID,
			Date:   dateParsed,
			TypeID: activities[0].TypeID,
			Value:  valueParsed,
		}
		err = daoFactory.ActivityDAO.UpdateActivity(activity)
		if err != nil {
			return errors.Wrap(err, "failed to save activity")
		}
		activities[0].Value = valueParsed
		return views.Render(ctx, components.Sum(activities[0]))
	})

	app.GET("/", func(ctx echo.Context) error {
		var dishes []model.MealDish
		// TODO: change this to be dynamic, need to create meals for dishes
		res := dbInstance.DB.Find(&dishes)
		if res.Error != nil {
			return res.Error
		}

		var activities []model.ActivityDetails
		res = dbInstance.
			DB.
			Table("activities").
			Joins("Join activity_types ON activities.type_id = activity_types.id").
			// TODO: change this to be dynamic, maybe remove from / page
			Where("activities.date = ?", "2024-02-02").
			Select("activities.*, activity_types.value_type as value_type, activity_types.name as name").
			Find(&activities)
		fmt.Printf("%+v", activities)
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
