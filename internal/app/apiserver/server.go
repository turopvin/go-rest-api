package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"github.com/turopvin/go-rest-api/internal/app/auth/api"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

func newServer(useCase auth.UseCase) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

	api.RegisterHttpEndPoints(s.router, useCase)
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
