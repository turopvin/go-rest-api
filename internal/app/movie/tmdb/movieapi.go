package tmdb

import (
	"github.com/turopvin/go-rest-api/internal/app/movie"
	"github.com/turopvin/go-rest-api/internal/app/movie/repository"
)

type MovieApi struct {
	movieRepository *repository.MovieRepository
}

func New(apiUrl, apiKey string) *MovieApi {
	return &MovieApi{
		movieRepository: repository.NewMovieRepository(apiUrl, apiKey),
	}
}

func (m *MovieApi) Movie() movie.Repositoru {
	if m.movieRepository != nil {
		return m.movieRepository
	}

	moviRep := &repository.MovieRepository{}
	return moviRep
}
