package api

import (
	"github.com/gorilla/mux"
	"github.com/turopvin/go-rest-api/internal/app/auth"
)

func RegisterHttpEndPoints(router *mux.Router, useCase auth.UseCase) {
	handler := New(useCase)

	router.HandleFunc("/auth/sign-up", handler.handleSignUp())
	router.HandleFunc("/auth/sign-in", handler.handleSignIn())
}
