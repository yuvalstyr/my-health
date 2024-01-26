package model

import "time"

type Activity struct {
	ID    string `gorm:"primaryKey"`
	Date  time.Time
	Type  string
	Count int
}
