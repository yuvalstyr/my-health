package model

type ActivityType struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	ValueType string
}
