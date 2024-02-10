package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"personal/health-app/service/daos"
	"personal/health-app/service/datebase"
	"personal/health-app/service/handlers"
	"personal/health-app/service/model"
	"personal/health-app/service/views"
	"personal/health-app/service/views/dashboard"
	"personal/health-app/service/views/dishes"
	"strings"

	"github.com/google/uuid"
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
			return errors.Wrap(err, "not activities found")
		}
		var activityTypes []model.ActivityType
		res = dbInstance.DB.Find(&activityTypes)
		if res.Error != nil {
			return res.Error
		}

		return views.Render(ctx, dashboard.Show(dishes, activities))
	})

	app.POST("/dish/add", func(ctx echo.Context) error {
		dish := ctx.FormValue("dish")
		dishScore := ctx.FormValue("meal_level")
		dishInput := model.MealDish{
			ID:     uuid.New().String(),
			Name:   dish,
			Score:  strings.ToLower(dishScore),
			MealID: "2",
		}
		res := dbInstance.DB.Create(&dishInput)
		if res.Error != nil {
			ctx.String(http.StatusInternalServerError, res.Error.Error())
			return res.Error
		}

		var dishesSlice []model.MealDish
		res = dbInstance.DB.Find(&dishesSlice)
		if res.Error != nil {
			ctx.String(http.StatusInternalServerError, res.Error.Error())
			return res.Error
		}

		err := views.Render(ctx, dishes.DishForm())
		if err != nil {
			ctx.String(http.StatusInternalServerError, res.Error.Error())
			return err
		}
		ctx.String(http.StatusOK, "dish created successfully")
		return views.Render(ctx, dishes.OOBDish(dishInput))
	})

	app.DELETE("/dish/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		res := dbInstance.DB.Delete(&model.MealDish{}, id)
		if res.Error != nil {
			return res.Error
		}

		var dishes []model.MealDish
		res = dbInstance.DB.Find(&dishes)
		if res.Error != nil {
			ctx.String(http.StatusInternalServerError, res.Error.Error())
			return res.Error
		}
		ctx.String(http.StatusNoContent, "dishes deleted successfully")
		return nil
	})

	app.PUT("/dish/:id", func(ctx echo.Context) error {
		db := dbInstance.DB
		id := ctx.Param("id")
		var dishInput model.MealDish
		res := db.First(dishInput, id)
		if res.Error != nil {
			errorMessage := fmt.Sprint("did't find dish with id ", id)
			ctx.String(http.StatusBadRequest, errorMessage)
			return res.Error
		}
		dish := ctx.FormValue("dish")
		dishLevel := ctx.FormValue("meal_level")
		dishInput.Name = dish
		dishInput.Score = strings.ToLower(dishLevel)
		res = db.Save(&dishInput)
		if res.Error != nil {
			ctx.String(http.StatusInternalServerError, "db error")
			return res.Error
		}
		ctx.String(http.StatusOK, "dish updated successfully")
		return views.Render(ctx, dishes.Dish(dishInput))
	})

	// when pressing the edit button on dish list screen, this endpoint is hit, and renders the form with the current dish data
	app.GET("/dish/:id/edit", func(ctx echo.Context) error {
		var dish model.MealDish
		id := ctx.Param("id")
		res := dbInstance.DB.First(&dish, id)
		if res.Error != nil {
			ctx.String(http.StatusInternalServerError, res.Error.Error())
			return res.Error
		}
		ctx.String(http.StatusOK, "dish updated successfully")
		return views.Render(ctx, dishes.DishFormRow(dish))
	})
	app.Start("localhost:4040")
}
