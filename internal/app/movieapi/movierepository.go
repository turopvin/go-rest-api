package movieapi

import "github.com/turopvin/go-rest-api/internal/app/model"

type MovieRepository interface {
	FindByTitle(title string) ([]model.Movie, error)
}
