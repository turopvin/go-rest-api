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
		title, err := h.useCase.FindByTitle("")
		if err != nil {

		}

	}
}
