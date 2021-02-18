package apiserver

import (
	"awesomeProject3/internal/app/model"
	"awesomeProject3/internal/app/store/teststore"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//This is Endpoint Tests

//Test Get endpoint
func TestServer_HandleGetStatistic(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/statistic/", nil)
	if err != nil {
		panic(err)
	}
	q := req.URL.Query()
	q.Add("from", "1999-02-13")
	q.Add("to", "2021-02-19")
	req.URL.RawQuery = q.Encode()
	s := NewServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}

//Test Post endpoint
func TestServer_HandleAddStatistic(t *testing.T) {
	rec := httptest.NewRecorder()
	body, err := json.Marshal(model.TestStatistic(t))
	if err != nil {
		panic(err)
	}
	str := strings.NewReader(string(body))
	req := httptest.NewRequest(http.MethodPost, "/statistic/", str)
	s := NewServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusCreated)
}

//Test Delete endpoint
func TestServer_HandleDeleteStatistic(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/statistic/", nil)
	s := NewServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
