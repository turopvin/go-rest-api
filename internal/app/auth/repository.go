package auth

import (
	"context"
	"github.com/turopvin/go-rest-api/internal/app/auth/model"
)

type UserRepository interface {
	CreateUser(context context.Context, user *model.User) error
	GetUser(context context.Context, username, password string) (*model.User, error)
	ParseToken()
}
