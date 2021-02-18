package sqlstore

import (
	"awesomeProject3/internal/app/model"
	"fmt"
	"time"
)

type StatisticRepository struct {
	store *Store
}

//Create ...
func (r *StatisticRepository) Create(u *model.Statistic) error {
	if err := u.Validate(); err != nil {
		return err
	}
	err := r.store.db.QueryRow(
		"INSERT INTO statistic (date, views, clicks, cost) VALUES ($1,$2,$3,$4)",
		u.Date,
		u.Views,
		u.Clicks,
		u.Cost)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (r *StatisticRepository) GetAll(FromTime time.Time, ToTime time.Time, SortRow string) ([]*model.StatisticWithMath, error) {
	if SortRow == "" {
		SortRow = "date"
	}
	QueryText := fmt.Sprintf(`
		SELECT date, views, clicks, cost,clicks/cost as cpc , views/cost  as cpm
		FROM statistic
		WHERE date >= $1 AND date <= $2
		ORDER BY %s ;`, SortRow)
	rows, err := r.store.db.Query(QueryText, FromTime, ToTime)
	if err != nil {
		return nil, err
	}
	us := make([]*model.StatisticWithMath, 0)
	for rows.Next() {
		u := new(model.StatisticWithMath)
		err := rows.Scan(&u.Date, &u.Views, &u.Clicks, &u.Cost, &u.Cpc, &u.Cpm)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return us, nil
}

func (r *StatisticRepository) Delete() error {
	_, err := r.store.db.Exec("TRUNCATE TABLE statistic;")
	if err != nil {
		return err
	}
	return nil
}
