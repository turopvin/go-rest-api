package api

import (
	"encoding/json"
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"net/http"
)

type Handler struct {
	useCase auth.UseCase
}

func NewUseCase(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n := new(signInput)
		if err := json.NewDecoder(r.Body).Decode(n); err != nil {
			sendError(w, r, http.StatusBadRequest, err)
			return
		}
		if err := h.useCase.SignUp(r.Context(), n.Username, n.Password); err != nil {
			sendError(w, r, http.StatusBadRequest, err)
		}
		sendRespond(w, r, http.StatusOK, nil)
	}
}

func (h *Handler) handleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n := new(signInput)
		if err := json.NewDecoder(r.Body).Decode(n); err != nil {
			sendError(w, r, http.StatusNotFound, err)
			return
		}
		token, err := h.useCase.SignIn(r.Context(), n.Username, n.Password)
		if err != nil {
			sendError(w, r, http.StatusNotFound, err)
		}
		sendRespond(w, r, http.StatusOK, token)
	}
}
