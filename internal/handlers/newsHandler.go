package handlers

import (
	"encoding/json"
	"net/http"
	"uroomBackend/internal/model"
)

func (s *server) handleCreateNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &ReqNews{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		news := &model.News{
			Date:             req.Date,
			Name:             req.Name,
			SmallDescription: req.SmallDescription,
			FullDescription:  req.FullDescription,
			ImagePath:        req.ImagePath,
		}

		if err := s.store.News().CreateNews(news); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, map[string]interface{}{"status": "success", "id": news.ID})
	}
}

func (s *server) handleNewsDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &ReqFindByID{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.News().DeleteNews(req.ID); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string]interface{}{"status": "delete success"})
	}
}

func (s *server) handleNewsFindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &ReqFindByID{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		news, err := s.store.News().FindById(req.ID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string]interface{}{"news": news})
	}
}

func (s *server) handleGetAllNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		news, err := s.store.News().GetAllNews()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string]interface{}{"news": news})
	}
}

func (s *server) handleNewsUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &ReqUpdateNews{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		updateNews := &model.News{
			ID:               req.ID,
			Date:             req.Date,
			Name:             req.Name,
			SmallDescription: req.SmallDescription,
			FullDescription:  req.FullDescription,
			ImagePath:        req.ImagePath,
		}

		if _, err := s.store.News().UpdateNews(updateNews); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string]interface{}{"status": "success"})
	}
}
