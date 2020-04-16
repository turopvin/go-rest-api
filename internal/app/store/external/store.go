package external

import (
	"github.com/turopvin/go-rest-api/internal/app/apiserver"
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"github.com/turopvin/go-rest-api/internal/app/movie"
	"github.com/turopvin/go-rest-api/internal/app/movie/repository"
)

type Store struct {
	movieRepository movie.Repository
}

func New(config *apiserver.Config) *Store {
	return &Store{}
}

func (s *Store) UserRepository() auth.UserRepository {
	panic("implement me")
}

func (s *Store) MovieRepository() movie.Repository {
	if s.movieRepository != nil {
		return s.movieRepository
	}

	mr := &repository.MovieRepository{}
}
