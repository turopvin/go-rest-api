package store

import (
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"github.com/turopvin/go-rest-api/internal/app/movie"
)

type Store interface {
	UserRepository() auth.UserRepository
	MovieRepository() movie.Repository
}
