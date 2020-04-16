package store

import "github.com/turopvin/go-rest-api/internal/app/auth"

type Store interface {
	UserRepository() auth.UserRepository
}
