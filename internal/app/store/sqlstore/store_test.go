package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("TEST_DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=statistic  sslmode= disable user=postgres password=yunis"
	}
	os.Exit(m.Run())
}
