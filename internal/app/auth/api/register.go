package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"net/http"
)

func RegisterHttpEndPoints(router *mux.Router, useCase auth.UseCase) {
	handler := New(useCase)

	router.HandleFunc("/auth/sign-up", handler.handleSignUp(context.Background())).Methods(http.MethodPost)
	router.HandleFunc("/auth/sign-in", handler.handleSignIn(context.Background())).Methods(http.MethodPost)
}
