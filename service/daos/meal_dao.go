package daos

import (
	"personal/health-app/service/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MealDAO struct {
	db *gorm.DB
}

func NewMealDAO(db *gorm.DB) MealDAOInterface {
	return &MealDAO{db: db}
}

type MealDAOInterface interface{}

func (d *MealDAO) GetMeals() ([]model.Meal, error) {
	var meals []model.Meal
	res := d.db.Find(&meals).Order("id asc")
	if res.Error != nil {
		return []model.Meal{}, errors.Wrap(res.Error, "Failed to get meals from db")
	}
	return meals, nil
}
