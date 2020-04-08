package store

import "github.com/turopvin/go-rest-api/internal/app/model"

type UserRepository interface {
	Create(u *model.User) error
	FindByEmail(s string) (*model.User, error)
}
