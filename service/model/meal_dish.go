package model

type MealDish struct {
	ID     string `gorm:"primaryKey;autoincrement"`
	Name   string
	Score  string
	MealID string
}
