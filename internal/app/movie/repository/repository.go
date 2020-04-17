package repository

import (
	"encoding/json"
	"errors"
	"github.com/turopvin/go-rest-api/internal/app/movie/model"
	"log"
	"net/http"
	"net/url"
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

type tmdbResponse struct {
	Page         int               `json:"page"`
	TotalResults int               `json:"total_results"`
	TotalPages   int               `json:"total_pages"`
	Results      []model.TmdbMovie `json:"results"`
}

type omdbResponse struct {
	Title string `json:"Title"`
	Year  string `json:"Year"`
}

type chanelStruct struct {
	ApiName string                `json:"api_name"`
	Movies  []model.ResponseMovie `json:"movies"`
}

func (m *MovieRepository) FindByTitle(title string) (map[string][]model.ResponseMovie, error) {
	resultMap := make(map[string][]model.ResponseMovie)

	channel := make(chan chanelStruct)
	errorChannel := make(chan error)
	go resultsTmdb(m.MovieApi.ApiTmdbUrl, m.MovieApi.ApiTmdbKey, title, channel, errorChannel)
	go resultsOmdb(m.MovieApi.ApiOmdbUrl, m.MovieApi.ApiOmdbKey, title, channel, errorChannel)

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

func resultsTmdb(apiUrl, apiKey, movieTitle string, channel chan<- chanelStruct, errorChannel chan<- error) {
	tmdbUrl, err := url.Parse(apiUrl)
	if err != nil {
		errorChannel <- err
		return
	}
	tmdbUrl.Path = "/3/search/movie"
	q := tmdbUrl.Query()
	q.Set("api_key", apiKey)
	q.Set("language", "en-US")
	q.Set("query", movieTitle)
	q.Set("include_adult", "false")
	tmdbUrl.RawQuery = q.Encode()

	resp, err := http.Get(tmdbUrl.String())
	if err != nil || resp.StatusCode != http.StatusOK {
		if err == nil {
			tmdbErr := errors.New("Request to Tmdb API failed")
			errorChannel <- tmdbErr
			log.Println(tmdbErr)
		}
		errorChannel <- err
		return
	}
	r := &tmdbResponse{}
	if err = json.NewDecoder(resp.Body).Decode(r); err != nil {
		errorChannel <- err
		return
	}
	movies := convertTmdbToResponseMovie(r.Results)

	channel <- chanelStruct{
		ApiName: "tmdb",
		Movies:  movies,
	}
}

func resultsOmdb(apiUrl, apiKey, movieTitle string, channel chan<- chanelStruct, errorChannel chan<- error) {
	omdbUrl, err := url.Parse(apiUrl)
	if err != nil {
		errorChannel <- err
		return
	}
	q := omdbUrl.Query()
	q.Set("apikey", apiKey)
	q.Set("t", movieTitle)
	omdbUrl.RawQuery = q.Encode()

	resp, err := http.Get(omdbUrl.String())
	if err != nil || resp.StatusCode != http.StatusOK {
		if err == nil {
			omdbErr := errors.New("Request to Omdb API failed")
			errorChannel <- omdbErr
			log.Println(omdbErr)
		}
		errorChannel <- err
		return
	}
	r := &omdbResponse{}
	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		errorChannel <- err
		return
	}
	omdbresult := model.ResponseMovie{
		Title:       r.Title,
		ReleaseDate: r.Year,
	}
	channel <- chanelStruct{
		ApiName: "omdb",
		Movies:  []model.ResponseMovie{omdbresult},
	}
}

func convertTmdbToResponseMovie(tmdbMovies []model.TmdbMovie) []model.ResponseMovie {
	var responseSlice []model.ResponseMovie
	for _, tmdb := range tmdbMovies {
		r := model.ResponseMovie{
			Title:       tmdb.Title,
			ReleaseDate: tmdb.ReleaseDate,
		}
		responseSlice = append(responseSlice, r)
	}
	return responseSlice
}
