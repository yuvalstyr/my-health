package model

import "time"

type ActivityDetails struct {
	Date      time.Time
	ID        string `gorm:"primaryKey"`
	TypeID    string
	Name      string
	ValueType string
	Value     int
}
