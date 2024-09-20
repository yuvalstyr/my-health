package model

type MealDish struct {
	ID     string `gorm:"primaryKey"`
	Name   string
	Score  string
	MealID string
}
