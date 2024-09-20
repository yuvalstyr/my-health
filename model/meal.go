package model

import "time"

type Meal struct {
	Date time.Time
	Id   string `gorm:"primaryKey"`
	chef Chef
	Type MealType
}
