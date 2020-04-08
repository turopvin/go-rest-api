package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/turopvin/go-rest-api/internal/app/model"
	"github.com/turopvin/go-rest-api/internal/app/store"
	"io"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s server) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/user/create", s.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/user/get", s.handleGetUser())
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

func (s *server) handleUserCreate() http.HandlerFunc {
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

		if err := s.store.User().Create(u); err != nil {
			s.logger.Debug("something happened")
		}
	}
}

func (s *server) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			s.logger.Debug("")
		}
		user, err := s.store.User().FindByEmail(email)
		if err != nil {
			io.WriteString(w, "No such user found")
		}
		json.NewEncoder(w).Encode(user)
	}
}
