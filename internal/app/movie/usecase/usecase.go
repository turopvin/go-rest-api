package usecase

import (
	"github.com/turopvin/go-rest-api/internal/app/movie"
	"github.com/turopvin/go-rest-api/internal/app/movie/model"
)

type MovieUseCase struct {
	repository movie.Repository
}

func New(repository movie.Repository) *MovieUseCase {
	return &MovieUseCase{repository: repository}
}

func (m *MovieUseCase) FindByTitle(title string) ([]model.ResponseMovie, error) {
	byTitle, err := m.repository.FindByTitle(title)
	if err != nil {
		return nil, err
	}

	return byTitle, nil

}
