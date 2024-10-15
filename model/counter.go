package model

type Counter struct {
	VisualData CounterVisualData
	Kpi        KPI       `gorm:"embedded"`
	Weekly     WeeklyKPI `gorm:"embedded"`
}
