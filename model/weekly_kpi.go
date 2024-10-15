package model

type WeeklyKPI struct {
	ID         string `gorm:"primaryKey"`
	WeekNumber string
	KPITypeId  string
	Value      int
}
