package sqlstore_test

import (
	"awesomeProject3/internal/app/model"
	"awesomeProject3/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStatisticRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, databaseURL)
	defer teardown("statistic")
	s := sqlstore.New(db)
	err := s.Statistic().Create(&model.Statistic{
		Date:   time.Now(),
		Views:  3,
		Clicks: 4,
		Cost:   4,
	})
	assert.NoError(t, err)
}

func TestStatisticRepository_GetAll(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, databaseURL)
	defer teardown("statistic")
	s := sqlstore.New(db)

	_, err := s.Statistic().GetAll(time.Now(), time.Now(), "")

	assert.NoError(t, err)
}

func TestStatisticRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDb(t, databaseURL)
	defer teardown("statistic")
	s := sqlstore.New(db)

	err := s.Statistic().Delete()

	assert.NoError(t, err)
}
