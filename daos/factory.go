package daos

import (
	"gorm.io/gorm"
)

type Factory struct {
	DishDAO DishDAOInterface
	Counter CounterDAOInterface
}

func NewDAOs(db *gorm.DB) *Factory {
	return &Factory{
		DishDAO: NewDishDAO(db),
		Counter: NewCounterDAO(),
	}
}
