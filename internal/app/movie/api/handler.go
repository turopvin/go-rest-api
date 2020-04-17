package api

import (
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
		title := r.URL.Query().Get("title")
		if len(title) < 1 {
			SendRespond(w, r, http.StatusBadRequest, nil)
		}
		movies, err := h.useCase.FindMoviesByTitle(title)
		if err != nil {
			SendError(w, r, http.StatusNotFound, nil)
		}
		SendRespond(w, r, http.StatusOK, movies)

	}
}
