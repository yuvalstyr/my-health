package datebase

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Instance struct {
	DB *gorm.DB
}

func New(config string) (*Instance, error) {
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Instance{
		DB: db,
	}, nil
}
