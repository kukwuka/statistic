package sqlstore

import (
	"awesomeProject3/internal/app/store"
	"database/sql"

	_ "github.com/lib/pq" // ...
)

//Store ...
type Store struct {
	db                  *sql.DB
	statisticRepository *StatisticRepository
}

//New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//Statistic ...
func (s *Store) Statistic() store.StatisticRepository {
	if s.statisticRepository != nil {
		return s.statisticRepository
	}
	s.statisticRepository = &StatisticRepository{
		store: s,
	}
	return s.statisticRepository
}
