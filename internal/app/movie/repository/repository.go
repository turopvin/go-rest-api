package repository

import (
	"github.com/turopvin/go-rest-api/internal/app/movie/model"
	"github.com/turopvin/go-rest-api/internal/app/movie/repository/omdb"
	"github.com/turopvin/go-rest-api/internal/app/movie/repository/tmdb"
)

type MovieRepository struct {
	MovieApi *movieApi
}

type movieApi struct {
	ApiTmdbUrl string
	ApiTmdbKey string
	ApiOmdbUrl string
	ApiOmdbKey string
}

func NewTmdb(tmdbUrl, tmdbKey, omdbUrl, omdbKey string) *movieApi {
	return &movieApi{
		ApiTmdbUrl: tmdbUrl,
		ApiTmdbKey: tmdbKey,
		ApiOmdbUrl: omdbUrl,
		ApiOmdbKey: omdbKey,
	}
}

func (m *MovieRepository) FindByTitle(title string) (map[string][]model.ResponseMovie, error) {
	resultMap := make(map[string][]model.ResponseMovie)

	channel := make(chan model.ChannelMovie)
	errorChannel := make(chan error)
	go tmdb.MovieByTitle(m.MovieApi.ApiTmdbUrl, m.MovieApi.ApiTmdbKey, title, channel, errorChannel)
	go omdb.MovieByTitle(m.MovieApi.ApiOmdbUrl, m.MovieApi.ApiOmdbKey, title, channel, errorChannel)

	for i := 0; i < 2; i++ {
		select {
		case result := <-channel:
			resultMap[result.ApiName] = result.Movies
		case errorResult := <-errorChannel:
			return nil, errorResult
		}
	}

	return resultMap, nil
}
