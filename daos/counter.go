package daos

import (
	datebase "personal/health-app/database"
	"personal/health-app/model"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
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

	var counters []*model.Counter
	return counters, db.
		Table("weekly_kpis").
		Joins("LEFT JOIN kpis ON weekly_kpis.kpi_type_id = kpis.id").
		Select("weekly_kpis.*, kpis.*").
		Where(model.WeeklyKPI{WeekNumber: strconv.Itoa(week)}).
		Find(&counters).
		Error
}

// UpdateCounter saves the given counter to the database. It assumes that the
// counter is already created and that the id is set. If the id is not set, it
// will panic.
func (cd *Counter) UpdateCounter(counter *model.Counter) error {
	err := cd.UpdateKPI(&counter.Kpi)
	if err != nil {
		return errors.Wrap(err, "Failed updating kpi")
	}
	err = cd.UpdateWeeklyKPI(&counter.Weekly)
	if err != nil {
		return errors.Wrap(err, "Failed updating weekly kpi")
	}
	log.Info("counter updated")
	return nil
}

func (cd *Counter) UpdateKPI(kpi *model.KPI) error {
	db, err := datebase.GetDB()
	if err != nil {
		return err
	}
	return db.Save(kpi).Error
}

func (cd *Counter) UpdateWeeklyKPI(weeklyKPI *model.WeeklyKPI) error {
	db, err := datebase.GetDB()
	if err != nil {
		return err
	}
	return db.Save(weeklyKPI).Error
}

func (cd *Counter) Get(id string) (*model.Counter, error) {
	var counter model.Counter
	kpi, err := cd.GetKPI(id)
	if err != nil {
		return nil, err
	}
	weekly, err := cd.GetWeeklyKPI(id)
	if err != nil {
		return nil, err
	}
	counter.Kpi = *kpi
	counter.Weekly = *weekly

	return &counter, nil
}
func (cd *Counter) GetKPI(id string) (*model.KPI, error) {
	db, err := datebase.GetDB()
	if err != nil {
		return nil, err
	}
	var kpi model.KPI
	return &kpi, db.First(&kpi, id).Error
}

func (cd *Counter) GetWeeklyKPI(id string) (*model.WeeklyKPI, error) {
	db, err := datebase.GetDB()
	if err != nil {
		return nil, err
	}
	var weeklyKPI model.WeeklyKPI
	return &weeklyKPI, db.First(&weeklyKPI, id).Error
}
