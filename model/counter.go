package model

type Counter struct {
	ID         string `gorm:"primaryKey"`
	WeekNumber string
	KPITypeId  string
	Value      int
}

type CounterEnriched struct {
	Counter
	Name      string
	Icon      string
	ValueType string
	Target    int
}
