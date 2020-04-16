package usecase

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"github.com/turopvin/go-rest-api/internal/app/auth/model"
	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *model.User `json:"user"`
}

type AuthUseCase struct {
	userRepository auth.UserRepository
	hashSalt       string
	signingKey     []byte
}

func New(repository auth.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepository: repository,
		hashSalt:       "hash_salt",
		signingKey:     []byte("signing_key"),
	}
}

func (a AuthUseCase) SignUp(ctx context.Context, username, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &model.User{
		Username: username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}
	if err := a.userRepository.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (a AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepository.GetUser(ctx, username, password)
	if err != nil {
		return "", err
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Second * 86400)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, errors.New("invalid access token")
}
