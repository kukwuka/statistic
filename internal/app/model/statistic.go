package model

import (
	"awesomeProject3/internal/app/apiserver/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

//Statistic ...
type Statistic struct {
	Date   time.Time `json:"date"`
	Views  int       `json:"views"`
	Clicks int       `json:"clicks"`
	Cost   float64   `json:"cost"`
}

type StatisticWithMath struct {
	Date   time.Time `json:"date"`
	Views  int       `json:"views"`
	Clicks int       `json:"clicks"`
	Cost   float64   `json:"cost"`
	Cpc    float64   `json:"cpc"`
	Cpm    float64   `json:"cpm"`
}
type StatisticWithoutTimeField struct {
	Date   string  `json:"date"`
	Views  int     `json:"views"`
	Clicks int     `json:"clicks"`
	Cost   float64 `json:"cost"`
}

//Validate ...
func (u *Statistic) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Date, validation.Required),
		validation.Field(&u.Views, validation.Required),
		validation.Field(&u.Clicks, validation.Required),
		validation.Field(&u.Cost, validation.Required),
	)
}

func (u *StatisticWithoutTimeField) ValidateForTest() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Date, validation.Required),
		validation.Field(&u.Views, validation.Required),
		validation.Field(&u.Clicks, validation.Required),
		validation.Field(&u.Cost, validation.Required),
	)
}

func (u *StatisticWithoutTimeField) ToStatistic() (*Statistic, error) {
	parsedTime, err := utils.ParseTime(u.Date)
	if err != nil {
		return nil, err
	}
	stat := &Statistic{
		Date:   parsedTime,
		Views:  u.Views,
		Clicks: u.Clicks,
		Cost:   u.Cost,
	}
	return stat, nil
}
