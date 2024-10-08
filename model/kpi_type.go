package model

type KPIType struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Icon      string
	ValueType string
	Target    int
}
