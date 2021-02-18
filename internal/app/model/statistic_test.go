package model_test

import (
	"awesomeProject3/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatistic_Validate(t *testing.T) {
	TestCases := []struct {
		name    string
		u       func() *model.StatisticWithoutTimeField
		isValid bool
	}{{
		name: "valid",
		u: func() *model.StatisticWithoutTimeField {
			return model.TestStatistic(t)
		},
		isValid: true,
	}, {
		name: "empty views",
		u: func() *model.StatisticWithoutTimeField {
			u := model.TestStatistic(t)
			u.Views = 0
			return u
		},
		isValid: false,
	}, {
		name: "empty views",
		u: func() *model.StatisticWithoutTimeField {
			u := model.TestStatistic(t)
			u.Views = 0
			return u
		},
		isValid: false,
	},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().ValidateForTest())
			} else {
				assert.Error(t, tc.u().ValidateForTest())
			}
		})
	}
	//u := model.TestStatistic(t)
	//assert.NoError(t, u.Validate())
}
