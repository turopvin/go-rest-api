package usecase

import "github.com/turopvin/go-rest-api/internal/app/movie"

type MovieUseCase struct {
	repository movie.Repository
}

func New(repository movie.Repository) *MovieUseCase {
	return &MovieUseCase{repository: repository}
}

func (m MovieUseCase) FindByName() {
	panic("implement me")
}
