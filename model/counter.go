package model

type Counter struct {
	ID          string `gorm:"primaryKey"`
	WeekNumber  string
	Icon        string
	Name        string
	kpi_type_id int
	Value       int
	Target      int
}
