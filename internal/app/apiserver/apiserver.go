package apiserver

import (
	"context"
	authUseCase "github.com/turopvin/go-rest-api/internal/app/auth/usecase"
	movieUseCase "github.com/turopvin/go-rest-api/internal/app/movie/usecase"
	"github.com/turopvin/go-rest-api/internal/app/store/external"
	"github.com/turopvin/go-rest-api/internal/app/store/mongostore"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func Start(config *Config) error {
	client, err := createDbClient(config)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	//create repository for auth domain and pass it to use case
	authStore := mongostore.New(client.Database("dev"))
	authUC := authUseCase.New(authStore.UserRepository())

	//create repository for movie domain and pass it to use case
	movieStore := external.New(config)
	movieUC := movieUseCase.New(movieStore.MovieRepository())

	srv := newServer(authUC, movieUC)
	return http.ListenAndServe(config.BindAddr, srv)
}

func createDbClient(config *Config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DatabaseURL))
	if err != nil {
		return nil, err
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return client, nil
}
