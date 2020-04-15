package repository

import "go.mongodb.org/mongo-driver/mongo"

//import (
//	"github.com/turopvin/go-rest-api/internal/app/store"
//	"go.mongodb.org/mongo-driver/mongo"
//)
//
//type UserRepository struct {
//	collection       *mongo.Client
//	userRepository *UserRepository
//}
//
//func New(collection *mongo.Client) *UserRepository {
//	return &UserRepository{
//		collection: collection,
//	}
//}
//
//func (s *UserRepository) User() store.UserRepository {
//	if s.userRepository != nil {
//		return s.userRepository
//	}
//
//	s.userRepository = &UserRepository{
//		store: s,
//	}
//
//	return s.userRepository
//}

type UserRepository struct {
	Collection *mongo.Collection
}

func (s UserRepository) CreateUser() {
	panic("implement me")
}

func (s UserRepository) GetUser() {
	panic("implement me")
}

func (s UserRepository) ParseToken() {
	panic("implement me")
}
