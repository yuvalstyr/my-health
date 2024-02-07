package daos

import (
	"personal/health-app/service/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ActivityDAO struct {
	db *gorm.DB
}

func NewActivityDAO(db *gorm.DB) ActivityDAOInterface {
	return &ActivityDAO{db: db}
}

type ActivityDAOInterface interface {
	GetActivityDetails(activityID string, date string) ([]model.ActivityDetails, error)
	UpdateActivity(activity *model.Activity) error
}

func (a *ActivityDAO) GetActivityDetails(activityID string, date string) ([]model.ActivityDetails, error) {
	conditions := make(map[string]interface{})
	if activityID != "" {
		conditions["activity_types.id"] = activityID
	}
	if date != "" {
		conditions["activities.date"] = date
	}
	var activities []model.ActivityDetails
	res := a.db.Table("activities").
		Joins("JOIN activity_types ON activity_types.id = activities.type_id").
		Where(conditions).
		Select("activities.*, activity_types.value_type as value_type, activity_types.name as name").
		Find(&activities)

	return activities, errors.Wrap(res.Error, "DB error when getting activity details")
}

func (a *ActivityDAO) UpdateActivity(activity *model.Activity) error {
	return a.db.Model(activity).Save(activity).Error
}
