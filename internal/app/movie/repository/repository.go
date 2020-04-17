package repository

import (
	"encoding/json"
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

	tmdb := make(chan chanelStruct)
	omdb := make(chan chanelStruct)
	go resultsTmdb(m.MovieApi.ApiTmdbUrl, m.MovieApi.ApiTmdbKey, title, tmdb)
	go resultsOmdb(m.MovieApi.ApiOmdbUrl, m.MovieApi.ApiOmdbKey, title, omdb)

	select {
	case s := <-tmdb:
		resultMap[s.ApiName] = s.Movies
	case s := <-omdb:
		resultMap[s.ApiName] = s.Movies
	}

	return resultMap, nil
}

func resultsTmdb(apiUrl, apiKey, movieTitle string, channel chan<- chanelStruct) {
	tmdbUrl, err := url.Parse(apiUrl)
	if err != nil {
		log.Fatal(err)
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
	if err != nil {
		log.Fatal(err)
		return
	}
	r := &tmdbResponse{}
	if err = json.NewDecoder(resp.Body).Decode(r); err != nil {
		log.Fatal(err)
		return
	}
	movies := convertTmdbToResponseMovie(r.Results)

	channel <- chanelStruct{
		ApiName: "tmdb",
		Movies:  movies,
	}
}

func resultsOmdb(apiUrl, apiKey, movieTitle string, channel chan<- chanelStruct) {
	omdbUrl, err := url.Parse(apiUrl)
	if err != nil {
		return
	}
	query := omdbUrl.Query()
	query.Set("apikey", apiKey)
	query.Set("t", movieTitle)
	omdbUrl.RawQuery = query.Encode()

	get, err := http.Get(omdbUrl.String())
	if err != nil {
		return
	}
	om := &omdbResponse{}
	if err := json.NewDecoder(get.Body).Decode(om); err != nil {
		return
	}
	omdbresult := model.ResponseMovie{
		Title:       om.Title,
		ReleaseDate: om.Year,
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
