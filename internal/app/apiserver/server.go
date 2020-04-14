package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/turopvin/go-rest-api/internal/app/store"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store, config *Config) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter(config)

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter(config *Config) {
	handler := newhandler(s.store, config, s.logger)

	s.router.HandleFunc("/hello", handler.handleHello())
	s.router.HandleFunc("/user/create", handler.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/user/get", handler.handleGetUser())
	s.router.HandleFunc("/movie/get", handler.handleMovieGet())
}
