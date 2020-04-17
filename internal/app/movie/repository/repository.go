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

func (m *MovieRepository) FindByTitle(title string) (map[string][]model.ResponseMovie, error) {
	resultMap := make(map[string][]model.ResponseMovie)
	//movieApi API
	tmdbUrl, err := url.Parse(m.MovieApi.ApiTmdbUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tmdbUrl.Path = "/3/search/movie"
	q := tmdbUrl.Query()
	q.Set("api_key", m.MovieApi.ApiTmdbKey)
	q.Set("language", "en-US")
	q.Set("query", title)
	q.Set("include_adult", "false")
	tmdbUrl.RawQuery = q.Encode()

	resp, err := http.Get(tmdbUrl.String())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	r := &tmdbResponse{}
	if err = json.NewDecoder(resp.Body).Decode(r); err != nil {
		log.Fatal(err)
		return nil, err
	}
	tmdbMovies := convertTmdbToResponseMovie(r.Results)
	resultMap["tmdbMovies"] = tmdbMovies
	//omdbAPI
	omdbUrl, err := url.Parse(m.MovieApi.ApiOmdbUrl)
	if err != nil {
		return nil, err
	}
	query := omdbUrl.Query()
	query.Set("apikey", m.MovieApi.ApiOmdbKey)
	query.Set("t", title)
	omdbUrl.RawQuery = query.Encode()

	get, err := http.Get(omdbUrl.String())
	if err != nil {
		return nil, err
	}
	om := &omdbResponse{}
	if err := json.NewDecoder(get.Body).Decode(om); err != nil {
		return nil, err
	}
	omdbresult := model.ResponseMovie{
		Title:       om.Title,
		ReleaseDate: om.Year,
	}
	resultMap["omdbresults"] = []model.ResponseMovie{omdbresult}
	return resultMap, nil
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

//func convertOmdbToResponseMovie(omdbMovies []omdbResponse) []model.ResponseMovie  {
//	var responseSlice []model.ResponseMovie
//	for _, omdb := range omdbMovies{
//		r := model.ResponseMovie{
//			Title:       omdb.Title,
//			ReleaseDate: omdb.Year,
//		}
//		responseSlice = append(responseSlice, r)
//	}
//	return responseSlice
//}
