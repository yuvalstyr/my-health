package model

type KPI struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Icon      string
	ValueType string
	Target    int
}
