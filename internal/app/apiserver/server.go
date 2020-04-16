package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/turopvin/go-rest-api/internal/app/auth"
	authApi "github.com/turopvin/go-rest-api/internal/app/auth/api"
	"github.com/turopvin/go-rest-api/internal/app/movie"
	movieApi "github.com/turopvin/go-rest-api/internal/app/movie/api"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

func newServer(authUseCase auth.UseCase, movieUseCase movie.UseCase) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

	//register authentication end-points
	//returns sub-router,
	//it is used for "authentication-required" functionality
	subRouter := authApi.RegisterHttpEndPoints(s.router, authUseCase)

	//register private("authentication-required") end-points
	movieApi.RegisterHttpEndpoints(subRouter, movieUseCase)
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
