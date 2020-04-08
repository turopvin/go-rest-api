package store

import (
	"context"
	"github.com/turopvin/go-rest-api/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	collection := r.store.db.Database("dev").Collection("users")
	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	filter := bson.D{{"email", email}}
	collection := r.store.db.Database("dev").Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
