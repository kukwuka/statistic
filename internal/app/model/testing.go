package model

import (
	"testing"
)

func TestStatistic(t *testing.T) *StatisticWithoutTimeField {
	return &StatisticWithoutTimeField{
		Date:   "2020-12-01",
		Views:  6,
		Clicks: 19,
		Cost:   5000,
	}
}
