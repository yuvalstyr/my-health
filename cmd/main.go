package main

import (
	"fmt"
	"log"
	"os"
	"personal/health-app/service/datebase"
	"personal/health-app/service/model"
	"personal/health-app/service/templates"
	"personal/health-app/service/views"
	"personal/health-app/service/views/activities"
	"personal/health-app/service/views/dishes"
	"strconv"
	"strings"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
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

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		var dishes []model.MealDish
		res := dbInstance.DB.Find(&dishes)
		if res.Error != nil {
			ctx.Status(500)
			return res.Error
		}
		var activities []model.Activity
		res = dbInstance.DB.Find(&activities)
		if res.Error != nil {
			ctx.Status(500)
			return res.Error
		}
		var activityTypes []model.ActivityType
		res = dbInstance.DB.Find(&activityTypes)
		if res.Error != nil {
			ctx.Status(fiber.StatusInternalServerError)
		}
		ctx.Status(200)
		return views.ComponentsHandler(ctx, templates.Page(dishes, activities))
	})

	app.Get("/dish/:id/edit", func(ctx *fiber.Ctx) error {
		var dish model.MealDish
		id := ctx.Params("id")
		res := dbInstance.DB.First(&dish, id)
		if res.Error != nil {
			ctx.Status(500)
			return res.Error
		}
		ctx.Status(200)
		return views.ComponentsHandler(ctx, dishes.DishFormRow(dish))
	})

	app.Post("/dish/add", func(ctx *fiber.Ctx) error {
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
			return res.Error
		}
		ctx.Status(200)
		var dishesSlice []model.MealDish
		res = dbInstance.DB.Find(&dishesSlice)
		if res.Error != nil {
			ctx.Status(500)
			return res.Error
		}

		err := views.ComponentsHandler(ctx, dishes.DishForm())
		if err != nil {
			return err
		}
		return views.ComponentsHandler(ctx, dishes.OOBDish(dishInput))
	})

	app.Delete("/dish/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		res := dbInstance.DB.Delete(&model.MealDish{}, id)
		if res.Error != nil {
			return res.Error
		}
		ctx.Status(fiber.StatusOK)
		var dishes []model.MealDish
		res = dbInstance.DB.Find(&dishes)
		if res.Error != nil {
			ctx.Status(fiber.StatusInternalServerError)
			return res.Error
		}
		ctx.Status(fiber.StatusNoContent)
		return nil
	})

	app.Put("/dish/:id", func(ctx *fiber.Ctx) error {
		db := dbInstance.DB
		id := ctx.Params("id")
		var dishInput model.MealDish
		res := db.First(dishInput, id)
		if res.Error != nil {
			errorMessage := fmt.Sprint("did't find dish with id ", id)
			ctx.Status(fiber.StatusBadRequest).SendString(errorMessage)
			return res.Error
		}
		dish := ctx.FormValue("dish")
		dishLevel := ctx.FormValue("meal_level")
		dishInput.Name = dish
		dishInput.Score = strings.ToLower(dishLevel)
		res = db.Save(&dishInput)
		if res.Error != nil {
			ctx.Status(fiber.StatusInternalServerError).SendString("db error")
			return res.Error
		}
		ctx.Status(fiber.StatusOK)
		return views.ComponentsHandler(ctx, dishes.Dish(dishInput))
	})

	app.Get("/activity/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var activity model.Activity
		res := dbInstance.DB.Model(&activity).Find(&activity, id)
		if res.Error != nil {
			ctx.Status(500)
		}
		ctx.Status(200)
		return nil
	})

	app.Post("/activity/add", func(ctx *fiber.Ctx) error {
		activityType := ctx.FormValue("type")
		count := ctx.FormValue("count")
		countInt, err := strconv.Atoi(count)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return errors.Wrap(err, "failed to add activity")
		}
		input := model.Activity{
			ID:    uuid.New().String(),
			Date:  time.Now(),
			Type:  activityType,
			Count: countInt,
		}
		res := dbInstance.DB.Create(&input)
		if res.Error != nil {
			return res.Error
		}
		ctx.Status(200)
		var activity model.Activity
		res = dbInstance.DB.First(&activity)
		if res.Error != nil {
			ctx.Status(500)
			return res.Error
		}

		var activityTypes []model.ActivityType
		res = dbInstance.DB.Find(&activityTypes)
		if res.Error != nil {
			ctx.Status(fiber.StatusInternalServerError)
		}
		err = views.ComponentsHandler(ctx, activities.ActivityForm())
		if err != nil {
			return err
		}
		return nil
	})

	err = app.Listen("localhost:4040")
	if err != nil {
		println(errors.Wrapf(err, "failed to start server").Error())
	}
}
