package auth

type UserRepository interface {
	CreateUser()
	GetUser()
	ParseToken()
}
