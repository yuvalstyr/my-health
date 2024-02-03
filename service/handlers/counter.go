package handlers

import (
	"net/http"
	"personal/health-app/service/model"
	"personal/health-app/service/views"
	"personal/health-app/service/views/components"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type counterHandler struct {
	DB *gorm.DB
}

type counterHandlerInterface interface {
	Increment(ctx echo.Context) error
	Decrement(ctx echo.Context) error
}

func newCounter(db *gorm.DB) *counterHandler {
	return &counterHandler{DB: db}
}

func (c *counterHandler) Increment(ctx echo.Context) error {
	return c.counter(ctx, "increment")
}

func (c *counterHandler) Decrement(ctx echo.Context) error {
	return c.counter(ctx, "decrement")
}

func (c *counterHandler) counter(ctx echo.Context, action string) error {
	counterID := ctx.Param("id")
	var result model.Activity
	res := c.DB.
		Joins("JOIN activity_types ON activity_types.id = activities.type_id").
		Where("activity_types.id = ?", counterID).
		First(&result)
	if res.Error != nil {
		return ctx.String(http.StatusBadRequest, res.Error.Error())
	}
	if action == "increment" {
		result.Count += 1
	} else {
		result.Count -= 1
	}
	res = c.DB.Save(&result)
	if res.Error != nil {
		return ctx.String(http.StatusBadRequest, res.Error.Error())
	}
	return views.Render(ctx, components.Counter(strconv.Itoa(result.Count)))
}
