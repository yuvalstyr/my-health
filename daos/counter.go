package daos

import (
	"fmt"
	datebase "personal/health-app/database"
	"personal/health-app/model"
)

type Counter struct{}

type CounterDAOInterface interface {
	GetCountersPerWeek(week int) ([]*model.Counter, error)
	UpdateCounter(counter *model.Counter) error
	Get(id string) (*model.Counter, error)
}

func NewCounterDAO() CounterDAOInterface {
	return &Counter{}
}

func (cd *Counter) GetCountersPerWeek(week int) ([]*model.Counter, error) {
	db, err := datebase.GetDB()
	if err != nil {
		return nil, err
	}
	var counter []*model.Counter
	return counter, db.
		Where(&model.Counter{WeekNumber: fmt.Sprint(week)}).
		Find(&counter).
		Error
}

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
