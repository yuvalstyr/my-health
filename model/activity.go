package model

import "time"

type Activity struct {
	ID     string `gorm:"primaryKey"`
	Date   time.Time
	TypeID string
	Value  int
}
