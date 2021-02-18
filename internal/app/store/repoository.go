package store

import (
	"awesomeProject3/internal/app/model"
	"time"
)

//StatisticRepository ...
type StatisticRepository interface {
	Create(statistic *model.Statistic) error
	GetAll(FromTime time.Time, ToTime time.Time, SortRow string) ([]*model.StatisticWithMath, error)
	Delete() error
}
