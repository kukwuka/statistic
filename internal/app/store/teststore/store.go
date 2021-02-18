package teststore

import (
	"awesomeProject3/internal/app/model"
	"awesomeProject3/internal/app/store"
)

//Store ...
type Store struct {
	statisticRepository *StatisticRepository
}

//New ...
func New() *Store {
	return &Store{}
}

//Statistic ...
func (s *Store) Statistic() store.StatisticRepository {
	if s.statisticRepository != nil {
		return s.statisticRepository
	}
	s.statisticRepository = &StatisticRepository{
		store:      s,
		statistics: make(map[int]*model.Statistic),
	}
	return s.statisticRepository
}
