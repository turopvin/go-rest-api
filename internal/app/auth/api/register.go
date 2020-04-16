package api

import (
	"github.com/gorilla/mux"
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"net/http"
)

func RegisterHttpEndPoints(router *mux.Router, useCase auth.UseCase) *mux.Router {
	handler := New(useCase)
	middleware := NewMiddleware(useCase)

	router.HandleFunc("/auth/sign-up", handler.handleSignUp()).Methods(http.MethodPost)
	router.HandleFunc("/auth/sign-in", handler.handleSignIn()).Methods(http.MethodPost)

	private := router.PathPrefix("/api").Subrouter()
	private.Use(middleware.AuthenticateUser)

}
