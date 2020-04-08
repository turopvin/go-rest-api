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

func (s *server) configureRouter() {

	handler := newhandler(s.store, s.logger)

	s.router.HandleFunc("/hello", handler.handleHello())
	s.router.HandleFunc("/user/create", handler.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/user/get", handler.handleGetUser())
}
