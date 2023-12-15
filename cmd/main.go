package main

import (
	fiberType "github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"personal/health-app/service/datebase"
	"personal/health-app/service/fiber"
	"personal/health-app/service/model"
	"personal/health-app/service/templates"
	"personal/health-app/service/views"
	"strings"
)

func main() {
	dbInstance, err := datebase.New("host=monorail.proxy.rlwy.net user=postgres password=34AB5gA5636443FE4Egc3-cGE-4*DC-G dbname=railway port=26753 sslmode=disable")
	if err != nil {
		println(errors.Wrapf(err, "failed to connect database").Error())
		panic(err)
	}

	app := fiber.New(dbInstance.DB)
	app.Get("/", func(ctx *fiberType.Ctx) error {
		var dishes []model.MealDish
		res := dbInstance.DB.Find(&dishes)
		if res.Error != nil {
			ctx.Status(500)
			return res.Error
		}
		ctx.Status(200)
		return views.ComponentsHandler(ctx, templates.Page(dishes))
	})

	app.Post("/add", func(ctx *fiberType.Ctx) error {
		dish := ctx.FormValue("dish")
		dishLevel := ctx.FormValue("meal_level")
		dishInput := model.MealDish{
			Name:   dish,
			Score:  strings.ToLower(dishLevel),
			MealID: "2",
		}
		res := dbInstance.DB.Create(&dishInput)
		if res.Error != nil {
			return res.Error
		}
		ctx.Status(200)

		var dishes []model.MealDish
		res = dbInstance.DB.Find(&dishes)
		if res.Error != nil {
			ctx.Status(500)
			return res.Error
		}

		err := views.ComponentsHandler(ctx, templates.DishForm())
		if err != nil {
			return err
		}
		return views.ComponentsHandler(ctx, templates.OOBDish(dishInput))
	})

	app.Delete("/dish/:id", func(ctx *fiberType.Ctx) error {
		id := ctx.Params("id")
		res := dbInstance.DB.Delete(&model.MealDish{}, id)
		if res.Error != nil {
			return res.Error
		}
		ctx.Status(200)
		var dishes []model.MealDish
		res = dbInstance.DB.Find(&dishes)
		if res.Error != nil {
			ctx.Status(500)
			return res.Error
		}

		return views.ComponentsHandler(ctx, templates.Page(dishes))
	})

	err = app.Listen("localhost:4040")
	if err != nil {
		println(errors.Wrapf(err, "failed to start server").Error())
	}

}
