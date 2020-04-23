package tmdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/turopvin/go-rest-api/internal/app/movie/model"
	"log"
	"net/http"
	"net/url"
	"time"
)

type tmdbMovieResponse struct {
	Page         int               `json:"page"`
	TotalResults int               `json:"total_results"`
	TotalPages   int               `json:"total_pages"`
	Results      []model.TmdbMovie `json:"results"`
}

type tmdbMovieVideosResponse struct {
	Id      int                     `json:"id"`
	Results []model.TmdbMovieVideos `json:"results"`
}

type videoLinkChannel struct {
	MovieId int
	Links   []string
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
	r := &tmdbMovieResponse{}
	if err = json.NewDecoder(resp.Body).Decode(r); err != nil {
		errorChannel <- err
		return
	}

	movieVideosChannel := make(chan videoLinkChannel)
	prepareVideoLinks(r, apiUrl, apiKey, movieVideosChannel, errorChannel)

	for movie := range movieVideosChannel {
		for i, m := range r.Results {
			if movie.MovieId == m.Id {
				r.Results[i].TrailerLinks = movie.Links
			}
		}
	}

	movies := convertTmdbToResponseMovie(r.Results)

	channel <- model.ChannelMovie{
		ApiName: "tmdb",
		Movies:  movies,
	}
}

func prepareVideoLinks(r *tmdbMovieResponse, apiUrl, apiKey string, movieVideosChannel chan<- videoLinkChannel, errorChannel chan<- error) {
	defer close(movieVideosChannel)

	tmbdMovieVideoUrl, _ := url.Parse(apiUrl)
	query := tmbdMovieVideoUrl.Query()
	query.Set("api_key", apiKey)
	tmbdMovieVideoUrl.RawQuery = query.Encode()

	for _, v := range r.Results {
		go func(movieId int, movieVideoChannel chan<- videoLinkChannel) {
			tmbdMovieVideoUrl.Path = fmt.Sprintf("/3/movie/%v/videos", movieId)
			response, err := http.Get(tmbdMovieVideoUrl.String())
			if err != nil || response.StatusCode != http.StatusOK {
				if err == nil {
					tmdbErr := errors.New("Request to Tmdb API failed")
					errorChannel <- tmdbErr
					log.Println(tmdbErr)
				}
				errorChannel <- err
				return
			}
			movieResp := &tmdbMovieVideosResponse{}
			if err := json.NewDecoder(response.Body).Decode(movieResp); err != nil {
				errorChannel <- err
			}

			var trailerLinks []string
			for _, v := range movieResp.Results {
				trailerLinks = append(trailerLinks, fmt.Sprintf("https://www.youtube.com/watch?v=%v", v.YoutubeKey))
			}

			movieVideosChannel <- videoLinkChannel{
				MovieId: movieId,
				Links:   trailerLinks,
			}
		}(v.Id, movieVideosChannel)
	}
	<-time.After(time.Second * 5)
	close(movieVideosChannel)
}

func convertTmdbToResponseMovie(tmdbMovies []model.TmdbMovie) []model.ResponseMovie {
	var responseSlice []model.ResponseMovie
	for _, tmdb := range tmdbMovies {
		r := model.ResponseMovie{
			Title:        tmdb.Title,
			ReleaseDate:  tmdb.ReleaseDate,
			Description:  tmdb.Overview,
			TrailerLinks: tmdb.TrailerLinks,
		}
		responseSlice = append(responseSlice, r)
	}
	return responseSlice
}
