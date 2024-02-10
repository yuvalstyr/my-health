package daos

import (
	"gorm.io/gorm"
)

type Factory struct {
	ActivityDAO ActivityDAOInterface
	DishDAO     DishDAOInterface
}

func NewDAOs(db *gorm.DB) *Factory {
	return &Factory{
		ActivityDAO: NewActivityDAO(db),
		DishDAO:     NewDishDAO(db),
	}
}
