package handlers

import (
	"fmt"
	"net/http"
	dao "personal/health-app/service/daos"
	"personal/health-app/service/model"
	"personal/health-app/service/views"
	"personal/health-app/service/views/dishes"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type dishHandler struct {
	dao dao.DishDAOInterface
}

type dishHandlerInterface interface {
	Create(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Update(ctx echo.Context) error
	Get(ctx echo.Context) error
}

func newDish(daoFactory dao.Factory) *dishHandler {
	return &dishHandler{dao: daoFactory.DishDAO}
}

func (h *dishHandler) Create(ctx echo.Context) error {
	dish := ctx.FormValue("dish")
	dishScore := ctx.FormValue("meal_level")
	// TODO: change this to be dynamic, need to create meals for dishes
	dishInput := model.MealDish{
		ID:     uuid.New().String(),
		Name:   dish,
		Score:  strings.ToLower(dishScore),
		MealID: "2",
	}

	err := h.dao.Create(dishInput)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "failed to create dish")
	}

	err = views.Render(ctx, dishes.DishForm())
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return err
	}
	ctx.String(http.StatusOK, "dish created successfully")
	return views.Render(ctx, dishes.OOBDish(dishInput))
}

func (h *dishHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err := h.dao.Delete(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "failed to delete dish")
		return err
	}

	ctx.String(http.StatusNoContent, "dishes deleted successfully")
	return nil
}

func (h *dishHandler) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	dishInput, err := h.dao.Get(id)
	if err != nil {
		errorMessage := fmt.Sprint("did't find dish with id ", id)
		ctx.String(http.StatusBadRequest, errorMessage)
		return err
	}
	dish := ctx.FormValue("dish")
	dishLevel := ctx.FormValue("meal_level")
	dishInput.Name = dish
	dishInput.Score = strings.ToLower(dishLevel)
	err = h.dao.Update(dishInput)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "db error")
		return err
	}
	ctx.String(http.StatusOK, "dish updated successfully")
	return views.Render(ctx, dishes.Dish(dishInput))
}

func (h *dishHandler) Get(ctx echo.Context) error {
	var dish model.MealDish
	id := ctx.Param("id")
	dish, err := h.dao.Get(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return err
	}
	ctx.String(http.StatusOK, "dish updated successfully")
	return views.Render(ctx, dishes.DishFormRow(dish))
}
