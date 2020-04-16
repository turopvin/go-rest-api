package movie

import "github.com/turopvin/go-rest-api/internal/app/movie/model"

type UseCase interface {
	FindByTitle(title string) ([]model.ResponseMovie, error)
}
