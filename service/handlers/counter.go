package handlers

import (
	"net/http"
	dao "personal/health-app/service/daos"
	"personal/health-app/service/model"
	"personal/health-app/service/views"
	"personal/health-app/service/views/components"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type counterHandler struct {
	dao dao.ActivityDAOInterface
}

type counterHandlerInterface interface {
	Increment(ctx echo.Context) error
	Decrement(ctx echo.Context) error
}

func newCounter(daoFactory dao.Factory) *counterHandler {
	return &counterHandler{dao: daoFactory.ActivityDAO}
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
	activities, err := c.dao.GetActivityDetails(counterID, date)
	if err != nil {
		return errors.Wrap(err, "failed to getting activities details")
	}
	if len(activities) > 1 {
		return errors.New("multiple activities found")
	}
	if action == "increment" {
		result.Value += 1
	} else {
		result.Value -= 1
	}
	activity := &model.Activity{
		ID:     counterID,
		Date:   dateParsed,
		TypeID: activities[0].TypeID,
		Value:  activities[0].Value,
	}
	err = c.dao.UpdateActivity(activity)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return views.Render(ctx, components.Counter(result))
}
