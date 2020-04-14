package tmdbapi

import (
	"encoding/json"
	"github.com/turopvin/go-rest-api/internal/app/model"
	"log"
	"net/http"
	"net/url"
)

type MovieRepository struct {
	apiurl   string `toml:"api_tmdb_base_url"`
	apikey   string `toml:"api_tmdb_key"`
	response response
}

type response struct {
	page         int           `json:"page"`
	totalresults int           `json:"total_results"`
	totalpages   int           `json:"total_pages"`
	results      []model.Movie `json:"results"`
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
	if err = json.NewDecoder(resp.Body).Decode(m.response); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return nil, nil
}
