package apiserver

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/turopvin/go-rest-api/internal/app/model"
	"github.com/turopvin/go-rest-api/internal/app/movieapi"
	"github.com/turopvin/go-rest-api/internal/app/movieapi/tmdbapi"
	"github.com/turopvin/go-rest-api/internal/app/store"
	"io"
	"net/http"
)

type handler struct {
	store    store.Store
	movieapi movieapi.MovieApi
	logger   *logrus.Logger
}

func newhandler(s store.Store, c *Config, l *logrus.Logger) *handler {
	return &handler{
		store:    s,
		movieapi: tmdbapi.New(c.ApiTmdbBaseUrl, c.ApiTmdbKey),
		logger:   l,
	}
}

func (h *handler) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

func (h *handler) handleUserCreate() http.HandlerFunc {
	type request struct {
		FullName string `json:"full_name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			FullName: req.FullName,
			Password: req.Password,
			Email:    req.Email,
		}

		if err := h.store.User().Create(u); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
		}

		u.Sanitize()
		h.respond(w, r, http.StatusCreated, u)
	}
}

func (h *handler) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			h.error(w, r, http.StatusNotFound, nil)
			return
		}
		user, err := h.store.User().FindByEmail(email)
		if err != nil {
			h.error(w, r, http.StatusNotFound, err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (h *handler) handleMovieGet() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := h.movieapi.Movie().FindByTitle("inception")
		if err != nil {
			h.error(w, r, http.StatusNotFound, err)
		}
	}
}

func (h *handler) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (h *handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
