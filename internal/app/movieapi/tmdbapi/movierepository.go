package tmdbapi

import (
	"encoding/json"
	"github.com/turopvin/go-rest-api/internal/app/model"
	"log"
	"net/http"
	"net/url"
)

type MovieRepository struct {
	apiurl string `toml:"api_tmdb_base_url"`
	apikey string `toml:"api_tmdb_key"`
}

type response struct {
	Page         int           `json:"page"`
	TotalResults int           `json:"total_results"`
	TotalPages   int           `json:"total_pages"`
	Results      []model.Movie `json:"results"`
}

func NewMovieRepository(apiUrl, apiKey string) *MovieRepository {
	return &MovieRepository{
		apiurl: apiUrl,
		apikey: apiKey,
	}
}

func (m *MovieRepository) FindByTitle(title string) ([]model.Movie, error) {
	u, err := url.Parse(m.apiurl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	u.Path = "/3/search/movie"
	q := u.Query()
	q.Set("api_key", m.apikey)
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

	var movies []model.Movie
	for _, v := range r.Results {
		i := append(movies, v)
		movies = i
	}

	return movies, nil
}
