package teststore_test

import (
	"awesomeProject3/internal/app/model"
	"awesomeProject3/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStatisticRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestStatistic(t)
	stat, err := u.ToStatistic()
	if err != nil {
		panic(err)
	}
	assert.NoError(t, s.Statistic().Create(stat))
}

func TestStatisticRepository_Delete(t *testing.T) {
	s := teststore.New()
	assert.NoError(t, s.Statistic().Delete())
}

func TestStatisticRepository_GetAll(t *testing.T) {
	s := teststore.New()
	_, err := s.Statistic().GetAll(time.Now(), time.Now(), "")
	assert.NoError(t, err)
}
