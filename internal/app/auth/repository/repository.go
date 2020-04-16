package repository

import (
	"context"
	"github.com/turopvin/go-rest-api/internal/app/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func (r UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	mongoUser := toMongoUser(user)
	_, err := r.Collection.InsertOne(ctx, mongoUser)
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) GetUser(ctx context.Context, username, password string) (*model.User, error) {
	user := new(User)
	err := r.Collection.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(user)

	if err != nil {
		return nil, err
	}

	return toModel(user), nil
}

func (r UserRepository) ParseToken() {
	panic("implement me")
}

func toMongoUser(u *model.User) *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
	}
}

func toModel(u *User) *model.User {
	return &model.User{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Password: u.Password,
	}
}
