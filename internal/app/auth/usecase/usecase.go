package usecase

import "github.com/turopvin/go-rest-api/internal/app/auth"

type AuthUseCase struct {
	userRepository auth.UserRepository
}

func New(repository auth.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepository: repository}
}

func (a AuthUseCase) SignUp() {
	panic("implement me")
}

func (a AuthUseCase) LogIn() {
	panic("implement me")
}
