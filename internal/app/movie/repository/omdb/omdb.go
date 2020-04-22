package omdb

import (
	"encoding/json"
	"errors"
	"github.com/turopvin/go-rest-api/internal/app/movie/model"
	"log"
	"net/http"
	"net/url"
)

func MovieByTitle(apiUrl, apiKey, movieTitle string, channel chan<- model.ChannelMovie, errorChannel chan<- error) {
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
	r := &model.OmdbMovie{}
	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		errorChannel <- err
		return
	}
	omdbresult := model.ResponseMovie{
		Title:       r.Title,
		ReleaseDate: r.Year,
		Director:    r.Director,
		ImdbRating:  r.ImdbRating,
		Actors:      r.Actors,
		Description: r.Plot,
	}
	channel <- model.ChannelMovie{
		ApiName: "omdb",
		Movies:  []model.ResponseMovie{omdbresult},
	}
}
