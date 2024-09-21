package model

type Counter struct {
	ID          string `gorm:"primaryKey"`
	WeekNumber  string
	kpi_type_id int
	Value       int
}
