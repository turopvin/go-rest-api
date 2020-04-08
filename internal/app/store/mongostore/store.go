package mongostore

import (
	"github.com/turopvin/go-rest-api/internal/app/store"
	"go.mongodb.org/mongo-driver/mongo"
)
import _ "github.com/go-sql-driver/mysql"

type Store struct {
	dbClient       *mongo.Client
	userRepository *UserRepository
}

func New(dbClient *mongo.Client) *Store {
	return &Store{
		dbClient: dbClient,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
