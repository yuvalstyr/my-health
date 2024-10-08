package daos

import (
	"fmt"
	datebase "personal/health-app/database"
	"personal/health-app/model"
)

type Counter struct{}

type CounterDAOInterface interface {
	GetCountersPerWeek(week int) ([]*model.CounterEnriched, error)
	UpdateCounter(counter *model.Counter) error
	Get(id string) (*model.Counter, error)
}

func NewCounterDAO() CounterDAOInterface {
	return &Counter{}
}

func (cd *Counter) GetCountersPerWeek(week int) ([]*model.CounterEnriched, error) {
	db, err := datebase.GetDB()
	if err != nil {
		return nil, err
	}

	var counters []*model.CounterEnriched
	return counters, db.
		Table("counters").
		Joins("LEFT JOIN kpi_types ON counters.kpi_type_id = kpi_types.id").
		Select("counters.*, kpi_types.*").
		Where(&model.Counter{WeekNumber: fmt.Sprint(week)}).
		Find(&counters).
		Error
}

// UpdateCounter saves the given counter to the database. It assumes that the
// counter is already created and that the id is set. If the id is not set, it
// will panic.
func (cd *Counter) UpdateCounter(counter *model.Counter) error {
	db, err := datebase.GetDB()
	if err != nil {
		return err
	}
	return db.Save(counter).Error
}

func (cd *Counter) Get(id string) (*model.Counter, error) {
	db, err := datebase.GetDB()
	if err != nil {
		return nil, err
	}
	var counter model.Counter
	return &counter, db.First(&counter, id).Error
}
