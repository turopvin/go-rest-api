package repository

import (
	"encoding/json"
	"github.com/turopvin/go-rest-api/internal/app/movie/model"
	"log"
	"net/http"
	"net/url"
)

type MovieRepository struct {
	Tmdb *tmdb
}

type tmdb struct {
	Apiurl string `toml:"api_tmdb_base_url"`
	Apikey string `toml:"api_tmdb_key"`
}

func NewTmdb(url, key string) *tmdb {
	return &tmdb{
		Apiurl: url,
		Apikey: key,
	}
}

type response struct {
	Page         int               `json:"page"`
	TotalResults int               `json:"total_results"`
	TotalPages   int               `json:"total_pages"`
	Results      []model.TmdbMovie `json:"results"`
}

func (m *MovieRepository) FindByTitle(title string) ([]model.ResponseMovie, error) {
	u, err := url.Parse(m.Tmdb.Apiurl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	u.Path = "/3/search/movie"
	q := u.Query()
	q.Set("api_key", m.Tmdb.Apikey)
	q.Set("language", "en-US")
	q.Set("query", title)
	q.Set("include_adult", "false")
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	r := &response{}
	if err = json.NewDecoder(resp.Body).Decode(r); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return convertToResponseMovie(r.Results), nil
}

func convertToResponseMovie(tmdbMovies []model.TmdbMovie) []model.ResponseMovie {
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
