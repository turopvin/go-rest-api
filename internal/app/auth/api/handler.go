package api

import (
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"net/http"
)

type Handler struct {
	useCase auth.UseCase
}

func New(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) handleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
