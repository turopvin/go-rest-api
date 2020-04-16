package api

import (
	"github.com/gorilla/mux"
	"github.com/turopvin/go-rest-api/internal/app/movie"
)

func RegisterHttpEndpoints(router *mux.Router, useCase movie.UseCase) {
	handler := New(useCase)

	router.HandleFunc("/movie/by-title", handler.findByTitle())
}
