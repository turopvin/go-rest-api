package external

import (
	"github.com/turopvin/go-rest-api/internal/app/auth"
	config2 "github.com/turopvin/go-rest-api/internal/app/config"
	"github.com/turopvin/go-rest-api/internal/app/movie"
	"github.com/turopvin/go-rest-api/internal/app/movie/repository"
)

type Store struct {
	movieRepository movie.Repository
	config          *config2.Config
}

func New(config *config2.Config) *Store {
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
