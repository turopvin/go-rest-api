package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

type Store struct {
	config         *Config
	db             *mongo.Client
	userRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	//db, err := sql.Open("mysql", s.config.DataBaseURL)
	client, err := mongo.NewClient(options.Client().ApplyURI(s.config.DataBaseURL))
	if err != nil {
		return nil
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return err
	}

	s.db = client
	return nil
}

func (s *Store) Close() {

}

func (s *Store) InitCollections() {
	database := s.db.Database("dev")
	database.Collection("users")
}

func (s *Store) UserService() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
