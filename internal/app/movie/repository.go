package movie

import "github.com/turopvin/go-rest-api/internal/app/model"

type Repository interface {
	FindByTitle(title string) ([]model.Movie, error)
}
