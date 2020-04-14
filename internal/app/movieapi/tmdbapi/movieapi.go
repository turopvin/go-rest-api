package tmdbapi

import (
	"github.com/turopvin/go-rest-api/internal/app/movieapi"
)

type MovieApi struct {
	movieRepository *MovieRepository
}

func New(apiUrl, apiKey string) *MovieApi {
	return &MovieApi{
		movieRepository: NewMovieRepository(apiUrl, apiKey),
	}
}

func (m *MovieApi) Movie() movieapi.MovieRepository {
	if m.movieRepository != nil {
		return m.movieRepository
	}

	moviRep := &MovieRepository{}
	return moviRep
}
