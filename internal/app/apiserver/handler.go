package apiserver

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/turopvin/go-rest-api/internal/app/model"
	"github.com/turopvin/go-rest-api/internal/app/store"
	"io"
	"net/http"
)

type handler struct {
	store  store.Store
	logger *logrus.Logger
}

func newhandler(s store.Store, l *logrus.Logger) *handler {
	return &handler{
		store:  s,
		logger: l,
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
			return
		}

		u := &model.User{
			FullName: req.FullName,
			Password: req.Password,
			Email:    req.Email,
		}

		if err := h.store.User().Create(u); err != nil {
			h.logger.Debug("something happened")
		}
	}
}

func (h *handler) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			h.logger.Debug("")
		}
		user, err := h.store.User().FindByEmail(email)
		if err != nil {
			io.WriteString(w, "No such user found")
		}
		json.NewEncoder(w).Encode(user)
	}
}
