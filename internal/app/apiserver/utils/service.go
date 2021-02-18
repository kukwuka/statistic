package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

const TimeLayout = "2006-01-02"

func ParseQueryTimeToFrom(r *http.Request) (fromTime time.Time, toTime time.Time, err error) {
	queryUrl := r.URL.Query()

	from := strings.Join(queryUrl["from"], "")
	to := strings.Join(queryUrl["to"], "")
	fromTime, err = ParseTime(from)
	if err != nil {
		return fromTime, toTime, errors.New("can not parse from date")
	}
	toTime, err = ParseTime(to)
	if err != nil {
		return fromTime, toTime, errors.New("can not parse to date")
	}
	return fromTime, toTime, err
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func SortRow(r *http.Request) (RowName string, err error) {

	PermittedRow := []string{"date", "views", "clicks", "cost", "cpc", "cpm"}
	queryUrl := r.URL.Query()
	ClientRow := strings.Join(queryUrl["sortby"], "")
	if ClientRow == "" {
		return "", err
	}
	if GetPermission := contains(PermittedRow, ClientRow); GetPermission {
		return ClientRow, nil
	}
	return ClientRow, errors.New("can't sort by row named:" + ClientRow)
}

func ParseTime(Time string) (ParsedTime time.Time, err error) {
	ParsedTime, err = time.Parse(TimeLayout, Time)
	return ParsedTime, err
}
