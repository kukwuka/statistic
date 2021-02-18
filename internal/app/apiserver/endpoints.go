package apiserver

import (
	"awesomeProject3/internal/app/apiserver/utils"
	"awesomeProject3/internal/app/model"
	"encoding/json"
	"net/http"
)

func (s *server) handleAddStatistic() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		s.logRequest(r)
		req := &model.StatisticWithoutTimeField{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		clientTime, err := utils.ParseTime(req.Date)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u := &model.Statistic{
			Date:   clientTime,
			Views:  req.Views,
			Clicks: req.Clicks,
			Cost:   req.Cost,
		}
		if err := s.store.Statistic().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleGetStatistic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logRequest(r)
		SortRow, err := utils.SortRow(r)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		fromTime, toTime, err := utils.ParseQueryTimeToFrom(r)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		result, err := s.store.Statistic().GetAll(fromTime, toTime, SortRow)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, result)
	}
}

func (s *server) handleDeleteStatistic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logRequest(r)
		err := s.store.Statistic().Delete()
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, "deleted")
	}
}
