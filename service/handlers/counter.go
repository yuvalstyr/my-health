package handlers

import (
	"net/http"
	"personal/health-app/service/model"
	"personal/health-app/service/views"
	"personal/health-app/service/views/components"
	"time"

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
	date := ctx.QueryParam("date")
	dateParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	var result model.ActivityDetails
	res := c.DB.
		Table("activities").
		Joins("JOIN activity_types ON activity_types.id = activities.type_id").
		Where("activity_types.id = ? AND activities.date = ?", counterID, date).
		Select("activities.*, activity_types.value_type as value_type, activity_types.name as name").
		First(&result)
	if res.Error != nil {
		return ctx.String(http.StatusBadRequest, res.Error.Error())
	}
	if action == "increment" {
		result.Value += 1
	} else {
		result.Value -= 1
	}
	res = c.DB.Save(&model.Activity{
		ID:     counterID,
		Date:   dateParsed,
		TypeID: result.TypeID,
		Value:  result.Value,
	})
	if res.Error != nil {
		return ctx.String(http.StatusBadRequest, res.Error.Error())
	}
	return views.Render(ctx, components.Counter(result))
}
