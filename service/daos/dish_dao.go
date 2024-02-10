package daos

import (
	"personal/health-app/service/model"

	"gorm.io/gorm"
)

type DishDAO struct {
	db *gorm.DB
}

func NewDishDAO(db *gorm.DB) DishDAOInterface {
	return &DishDAO{db: db}
}

type DishDAOInterface interface {
	Create(dish model.MealDish) error
	Delete(id string) error
	Get(id string) (model.MealDish, error)
	Update(dish model.MealDish) error
}

func (d DishDAO) Create(dish model.MealDish) error {
	return d.db.Create(&dish).Error
}

func (d DishDAO) Delete(id string) error {
	return d.db.Delete(&model.MealDish{}, id).Error
}

func (d DishDAO) Get(id string) (model.MealDish, error) {
	var dish model.MealDish
	return dish, d.db.First(&dish, id).Error
}

func (d DishDAO) Update(dish model.MealDish) error {
	return d.db.Save(&dish).Error
}
