package api

import (
	"github.com/turopvin/go-rest-api/internal/app/apiserver"
	"github.com/turopvin/go-rest-api/internal/app/movie"
	"net/http"
)

type Handler struct {
	useCase movie.UseCase
}

func New(useCase movie.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) findByTitle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := h.useCase.FindByTitle("")
		if err != nil {
			apiserver.SendError(w, r, http.StatusNotFound, nil)
		}
		apiserver.SendRespond(w, r, http.StatusOK, movies)

	}
}
