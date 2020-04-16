package api

import (
	"context"
	"encoding/json"
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

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) handleSignUp(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n := new(signInput)
		if err := json.NewDecoder(r.Body).Decode(n); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := h.useCase.SignUp(ctx, n.Username, n.Password); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
		}
	}
}

func (h *Handler) handleSignIn(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n := new(signInput)
		if err := json.NewDecoder(r.Body).Decode(n); err != nil {
			h.error(w, r, http.StatusNotFound, err)
			return
		}

		if err := h.useCase.SignIn(ctx, n.Username, n.Password); err != nil {
			h.error(w, r, http.StatusNotFound, err)
			return
		}
	}
}

func (h *Handler) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (h *Handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
