package model

import "time"

type ActivityType struct {
	Date time.Time
	ID   string `gorm:"primaryKey"`
	Name string
}
