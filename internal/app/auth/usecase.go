package auth

import (
	"context"
	"github.com/turopvin/go-rest-api/internal/app/auth/model"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*model.User, error)
}
