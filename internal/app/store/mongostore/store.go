package mongostore

import (
	"github.com/turopvin/go-rest-api/internal/app/auth"
	"github.com/turopvin/go-rest-api/internal/app/auth/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db             *mongo.Database
	userRepository auth.UserRepository
}

func New(db *mongo.Database) *Store {
	return &Store{
		db: db,
	}
}

func (t *Store) UserRepository() auth.UserRepository {
	if t.userRepository != nil {
		return t.userRepository
	}

	ur := &repository.UserRepository{
		Collection: t.db.Collection("users"),
	}
	t.userRepository = ur
	return ur
}
