package movie

import "github.com/turopvin/go-rest-api/internal/app/movie/model"

type UseCase interface {
	FindMoviesByTitle(title string) (map[string][]model.ResponseMovie, error)
}
