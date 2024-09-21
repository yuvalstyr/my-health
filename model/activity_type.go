package model

type KPIType struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	ValueType string
}
