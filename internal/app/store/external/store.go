package external

import (
	"github.com/turopvin/go-rest-api/internal/app/apiserver"
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"github.com/turopvin/go-rest-api/internal/app/movie"
	"github.com/turopvin/go-rest-api/internal/app/movie/repository"
)

type Store struct {
	movieRepository movie.Repository
	config          *apiserver.Config
}

func New(config *apiserver.Config) *Store {
	return &Store{config: config}
}

func (s *Store) UserRepository() auth.UserRepository {
	panic("implement me")
}

func (s *Store) MovieRepository() movie.Repository {
	if s.movieRepository != nil {
		return s.movieRepository
	}
	t := repository.NewTmdb(s.config.ApiTmdbBaseUrl, s.config.ApiTmdbKey)
	mr := &repository.MovieRepository{Tmdb: t}
	s.movieRepository = mr
	return s.movieRepository
}
