package movie

import (
	"github.com/turopvin/go-rest-api/internal/app/movie/model"
)

type Repository interface {
	FindByTitle(title string) (map[string][]model.ResponseMovie, error)
}
