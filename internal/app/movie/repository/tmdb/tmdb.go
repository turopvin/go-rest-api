package tmdb

import (
	"encoding/json"
	"errors"
	"github.com/turopvin/go-rest-api/internal/app/movie/model"
	"log"
	"net/http"
	"net/url"
)

type tmdbResponse struct {
	Page         int               `json:"page"`
	TotalResults int               `json:"total_results"`
	TotalPages   int               `json:"total_pages"`
	Results      []model.TmdbMovie `json:"results"`
}

func MovieByTitle(apiUrl, apiKey, movieTitle string, channel chan<- model.ChannelMovie, errorChannel chan<- error) {
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

	channel <- model.ChannelMovie{
		ApiName: "tmdb",
		Movies:  movies,
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
