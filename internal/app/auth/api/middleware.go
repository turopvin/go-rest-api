package api

import (
	"context"
	"errors"
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"net/http"
	"strings"
)

const (
	ctxKeyUser ctxKey = iota
)

type ctxKey int8

type Middleware struct {
	useCase auth.UseCase
}

func NewsMiddleware(usecase auth.UseCase) *Middleware {
	return &Middleware{
		useCase: usecase,
	}
}

func (m *Middleware) AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			sendError(w, r, http.StatusUnauthorized, nil)
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			sendError(w, r, http.StatusUnauthorized, nil)
			return
		}

		if headerParts[0] != "Bearer" {
			sendError(w, r, http.StatusUnauthorized, nil)
			return
		}

		user, err := m.useCase.ParseToken(context.Background(), headerParts[1])
		if err != nil {
			status := http.StatusInternalServerError
			if err == errors.New("invalid access token") {
				status = http.StatusUnauthorized
			}

			sendError(w, r, status, nil)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, user)))
	})
}
