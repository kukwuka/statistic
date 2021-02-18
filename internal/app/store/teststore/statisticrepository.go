package teststore

import (
	"awesomeProject3/internal/app/model"
	"time"
)

//For TestStore Unnecessary

//StatisticRepository ...
type StatisticRepository struct {
	store      *Store
	statistics map[int]*model.Statistic
}

//Create ...
func (r *StatisticRepository) Create(u *model.Statistic) error {
	if err := u.Validate(); err != nil {
		return err
	}

	r.statistics[u.Views] = u
	return nil
}

func (r *StatisticRepository) GetAll(FromTime time.Time, ToTime time.Time, SortRow string) ([]*model.StatisticWithMath, error) {
	return nil, nil
}

func (r *StatisticRepository) Delete() error {
	return nil
}
